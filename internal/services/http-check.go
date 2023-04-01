package services

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const DefaultRequestTimeout = 5 * time.Second

type HttpCheckConfig struct {
	Sender  *Sender
	URL     string
	Timeout time.Duration
}

// New creates new HTTP service health check that verifies the following:
// - connection establishing
// - getting response status from defined URL
// - verifying that status code is above 400
func NewHttpCheck(cfg *HttpCheckConfig) func(ctx context.Context) error {
	if cfg.Timeout == 0 {
		cfg.Timeout = DefaultRequestTimeout
	}

	return func(ctx context.Context) error {
		req, err := http.NewRequest(http.MethodGet, cfg.URL, nil)
		if err != nil {
			log.Errorf("[HttpCheck] Creating the request for the health check failed: %w", err)
			return nil
		}

		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
		defer cancel()

		// Inform remote service to close the connection after the transaction is complete
		req.Header.Set("Connection", "close")
		req = req.WithContext(ctx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Errorf("[HttpCheck] Making the request for the health check failed: %w", err)
			return nil
		}
		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			if cfg.Sender != nil {
				cfg.Sender.Send(SenderData{
					to:   "",
					text: "Service <code>" + resp.Status + "</code> is down\n" + "Status: <b>" + cfg.URL + "</b>",
				})
			}
			log.Errorf("[HttpCheck] %v service is not available at the moment", cfg.URL)
		}

		return nil
	}
}
