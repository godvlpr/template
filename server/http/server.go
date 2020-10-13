package http

import (
	"crypto/tls"
	"fmt"
	"github.com/godvlpr/template/config"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	log *logrus.Entry
	cfg config.Config
}

func NewServer() *Server {
	return &Server{
		log: config.GetConfig().GetLogCopy().Entry,
		cfg: config.GetConfig(),
	}
}

func (s Server) Listen(router *echo.Echo) error {

	serverHost := fmt.Sprintf("%s:%s", s.cfg.GetHTTPCopy().Host, s.cfg.GetHTTPCopy().Port)
	s.log.WithField("api", "start").
		Info(fmt.Sprintf("listenig addr =  %s, tls = %v", serverHost, s.cfg.GetHTTPCopy().SSL))

	httpServer := http.Server{
		Addr:           serverHost,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	switch s.cfg.GetHTTPCopy().SSL {
	case true:
		tlsConfig := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

				// Best disabled, as they don't provide Forward Secrecy,
				// but might be necessary for some clients
				// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			},
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519, // Go 1.8 only
			},
			InsecureSkipVerify: true,
		}

		httpServer.TLSConfig = tlsConfig
		if err := httpServer.ListenAndServeTLS(s.cfg.GetHTTPCopy().ServerCertPath, s.cfg.GetHTTPCopy().ServerKeyPath); err != nil {
			return errors.Wrap(err, "failed to start https server")
		}

	default:
		if err := httpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "failed to start http server")
		}
	}

	return nil
}
