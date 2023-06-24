package services

import (
	"context"
	"net/http"
	"time"

	adapters "github.com/WildEgor/gChecker/internal/adapters"
	"github.com/WildEgor/gChecker/internal/config"
	log "github.com/sirupsen/logrus"
)

// Service Facade contains logic for check services and send notifications
type CheckerService struct {
	telegramAdapter    adapters.ITelegramAdapter
	healthCheckAdapter *adapters.HealthCheckAdapter
	servicesConfig     *config.ServicesConfig
}

func NewCheckerService(
	telegramAdapter adapters.ITelegramAdapter,
	healthCheckAdapter *adapters.HealthCheckAdapter,
	servicesConfig *config.ServicesConfig,
) *CheckerService {
	return &CheckerService{
		telegramAdapter:    telegramAdapter,
		healthCheckAdapter: healthCheckAdapter,
		servicesConfig:     servicesConfig,
	}
}

// HINT: simple running in goroutine
func (s *CheckerService) Check() {
	log.Info(s.servicesConfig.URLs)
	for {
		sleep := time.Duration(s.servicesConfig.Timeout)
		for i := 0; i < len(s.servicesConfig.URLs); i++ {
			if !s.servicesConfig.URLs[i].Enabled {
				continue
			}

			func() {
				log.Info("Start check...")
				resp, err := http.Get(s.servicesConfig.URLs[i].URL)
				if err != nil {
					log.Warn("Error connect to server: " + s.servicesConfig.URLs[i].URL)
					return
				} else {
					log.Debug("Connect to server: " + s.servicesConfig.URLs[i].URL + " OK!")
				}
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					if err := s.telegramAdapter.Send("", "Service <code>"+resp.Status+"</code> is down\n"+"Status: <b>"+s.servicesConfig.URLs[i].URL+"</b>"); err != nil {
						log.Warn("Cannot send telegram alert. Reason: ", err)
					}
					sleep = time.Duration(s.servicesConfig.Timeout)
				}
			}()
		}

		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

// HINT: more advanced use case
func (s *CheckerService) ServicesCheck(ctx context.Context) {
	for i := 0; i < len(s.servicesConfig.URLs); i++ {
		if !s.servicesConfig.URLs[i].Enabled {
			continue
		}

		log.Info("Checking service: ", s.servicesConfig.URLs[i].URL)

		s.healthCheckAdapter.Register(adapters.HealthConfig{
			Name:      s.servicesConfig.URLs[i].ID,
			Timeout:   time.Duration(s.servicesConfig.Timeout),
			SkipOnErr: false,
			Check: NewHttpCheck(&HttpCheckConfig{
				Sender: InitSender(s.telegramAdapter),
				URL:    s.servicesConfig.URLs[i].URL,
			}),
		})
	}

	for {
		s.healthCheckAdapter.Measure(ctx)
		time.Sleep(time.Duration(5) * time.Second)
	}
}
