package storage

import (
	"github.com/godvlpr/template/storage/db"
)

type StorageProvider struct {
	DB db.QInterface
}

func New() *StorageProvider {
	return &StorageProvider{
		DB: db.New(),
	}
}
