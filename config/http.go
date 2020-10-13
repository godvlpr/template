package config

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/caarlos0/env"
)

type HTTP struct {
	Host           string `env:"TEMPLATE_API_HOST" envDefault:"localhost"`
	Port           string `env:"TEMPLATE_API_PORT"  envDefault:"1123"`
	SSL            bool   `env:"TEMPLATE_API_SSL" envDefault:"false"`
	ServerCertPath string `env:"TEMPLATE_API_CERT_PATH" envDefault:""`
	ServerKeyPath  string `env:"TEMPLATE_API_CERT_KEY" envDefault:""`
}

func (c Cfg) GetHTTPCopy() *HTTP {
	return &HTTP{
		Host:           c.HTTP.Host,
		Port:           c.HTTP.Port,
		SSL:            c.HTTP.SSL,
		ServerCertPath: c.HTTP.ServerCertPath,
		ServerKeyPath:  c.HTTP.ServerKeyPath,
	}
}

func (h HTTP) URL() (*url.URL, error) {
	if h.SSL {
		resultURL, err := url.Parse(fmt.Sprintf("https://%s:%s", h.Host, h.Port))
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse url")
		}

		return resultURL, nil
	}

	resultURL, err := url.Parse(fmt.Sprintf("http://%s:%s", h.Host, h.Port))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	return resultURL, nil
}

func  (c Cfg) LoadHTTP() *HTTP {
	http := &HTTP{}
	if err := env.Parse(http); err != nil {
		panic(err)
	}

	return http
}
