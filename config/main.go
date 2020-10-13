package config

import (
	"sync"
)

var (
	doOnce sync.Once
	config *Cfg
)

type Config interface {
	GetS3Copy() *S3
	GetLogCopy() *Log
	GetHTTPCopy() *HTTP
	GetAuthCopy() *Authentication
	GetDBCopy() *DBConf
}

type Cfg struct {
	S3   *S3
	Log  *Log
	HTTP *HTTP
	Auth *Authentication
	DB   *DBConf
}

func GetConfig() Config {
	doOnce.Do(load)
	return Cfg{
		S3:   config.GetS3Copy(),
		Log:  config.GetLogCopy(),
		HTTP: config.GetHTTPCopy(),
		Auth: config.GetAuthCopy(),
		DB:   config.GetDBCopy(),
	}
}

func load() {
	config = &Cfg{}

	config.S3 = config.LoadS3()
	config.Log = config.LoadLog()
	config.HTTP = config.LoadHTTP()
	config.Auth = config.LoadAuthentication()
	config.DB = config.LoadDB()
}
