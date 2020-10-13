package db

import (
	"github.com/godvlpr/template/models"
)

type UsersQ interface {
	Create(user models.User) error
}

type UsersWrapper struct {
	parent *DB
}

func (d *DB) UsersQ() UsersQ {
	return &UsersWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (pi UsersWrapper) Create(user models.User) error {
	return pi.parent.db.Model(&user).Insert()
}
