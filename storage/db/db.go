package db

import dbx "github.com/go-ozzo/ozzo-dbx"

type QInterface interface {
	DBX() *dbx.DB
	UsersQ() UsersQ
}
