package db

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/godvlpr/template/config"
	_ "github.com/lib/pq"
)

type DB struct {
	db *dbx.DB
}

func (d DB) DBX() *dbx.DB {
	return d.db
}

func New() QInterface {
	cfg := config.GetConfig()

	db, err := dbx.Open("postgres", cfg.GetDBCopy().Info())
	if err != nil {
		cfg.GetLogCopy().WithError(err).Fatal("failed to connect to db")
	}

	return &DB{
		db: db,
	}
}
