package config

import (
	"github.com/caarlos0/env"
)

type S3 struct {
	AccessKey string `env:"TEMPLATE_USER_MS_S3_ACCESS_KEY" envDefault:"qwe111"`
	SecretKey string `env:"TEMPLATE_USER_MS_S3_SECRET_KEY" envDefault:"qwe222"`
	Endpoint  string `env:"TEMPLATE_USER_MS_S3_ENDPOINT" envDefault:"qwe333"`
}

func (c Cfg) LoadS3() *S3 {
	s3 := &S3{}
	if err := env.Parse(s3); err != nil {
		panic(err)
	}

	return s3
}

func (c Cfg) GetS3Copy() *S3 {
	return &S3{
		AccessKey: c.S3.AccessKey,
		SecretKey: c.S3.SecretKey,
		Endpoint:  c.S3.Endpoint,
	}
}
