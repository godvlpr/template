package config

import (
	"github.com/caarlos0/env"
)

type Authentication struct {
	VerifyKey string `env:"TEMPLATE_AUTHENTICATION_SECRET" envDefault:"qwertyqwerty123"`
	Algorithm string `env:"TEMPLATE_AUTHENTICATION_ALGORITHM" envDefault:"HS256"`
}

func (c Cfg) GetAuthCopy() *Authentication {
	return &Authentication{
		VerifyKey: c.Auth.VerifyKey,
		Algorithm: c.Auth.Algorithm,
	}
}

func (c Cfg) LoadAuthentication() *Authentication {
	jwt := &Authentication{}
	if err := env.Parse(jwt); err != nil {
		panic(err)
	}

	return jwt
}
