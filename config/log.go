package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Log struct {
	*logrus.Entry
	Lvl string `env:"TEMPLATE_LOG_LEVEL" envDefault:"debug"`
}

func  (c Cfg) LoadLog() *Log {
	log := &Log{}

	if err := env.Parse(log); err != nil {
		panic(err)
	}

	logger := logrus.New()
	level, _ := logrus.ParseLevel(log.Lvl)
	logger.SetLevel(level)
	log.Entry = logrus.NewEntry(logger)

	return log
}

func (c Cfg) GetLogCopy() *Log {
	return &Log{
		Entry: c.Log.Entry,
		Lvl:   c.Log.Lvl,
	}
}
