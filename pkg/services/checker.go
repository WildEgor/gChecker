package services

import (
	"net/http"
	"time"

	adapters "github.com/WildEgor/checker/pkg/adapters"
	"github.com/WildEgor/checker/pkg/config"
	log "github.com/sirupsen/logrus"
)

type CheckerService struct {
	telegramAdapter adapters.ITelegramAdapter
	servicesConfig  *config.ServicesConfig
}

func NewCheckerService(
	telegramAdapter adapters.ITelegramAdapter,
	servicesConfig *config.ServicesConfig,
) *CheckerService {
	return &CheckerService{
		telegramAdapter: telegramAdapter,
		servicesConfig:  servicesConfig,
	}
}

// HINT: running in goroutine
func (s *CheckerService) Check() {
	log.Info(s.servicesConfig.ServiceUrls)
	for {
		sleep := time.Duration(s.servicesConfig.CheckTimeout)
		for _, service := range s.servicesConfig.ServiceUrls {
			func() {
				log.Info("Start check...")
				resp, err := http.Get(service.URL)
				if err != nil {
					log.Warn("Error connect to server: " + service.URL)
					return
				} else {
					log.Debug("Connect to server: " + service.URL + " OK!")
				}
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					s.telegramAdapter.SendAlert(service.URL, resp.Status)
					sleep = time.Duration(s.servicesConfig.MessTimeout)
				}
			}()
		}
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}
