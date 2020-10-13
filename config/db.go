package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

// DB is store all info about database connection
type DBConf struct {
	Name     string `env:"TEMPLATE_DATABASE_NAME" envDefault:"template"`
	Host     string `env:"TEMPLATE_DATABASE_HOST" envDefault:"localhost"`
	Port     int    `env:"TEMPLATE_DATABASE_PORT" envDefault:"5444"`
	User     string `env:"TEMPLATE_DATABASE_USER" envDefault:"postgres"`
	Password string `env:"TEMPLATE_DATABASE_PASSWORD" envDefault:"1234567"`
	SSL      string `env:"TEMPLATE_DATABASE_SSL" envDefault:"disable"`
}

func (c Cfg) GetDBCopy() *DBConf {
	return &DBConf{
		Name:     c.DB.Name,
		Host:     c.DB.Host,
		Port:     c.DB.Port,
		User:     c.DB.User,
		Password: c.DB.Password,
		SSL:      c.DB.SSL,
	}
}

// Info returning compiled statement. Statement is uri string which is provide access to the db
func (d DBConf) Info() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSL,
	)
}

// DB creates new local conn of the database
func  (c Cfg) LoadDB() *DBConf {
	database := &DBConf{}
	if err := env.Parse(database); err != nil {
		panic(err)
	}

	return database
}