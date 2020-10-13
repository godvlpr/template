package handlers

import (
	"github.com/godvlpr/template/server/http/handlers/userhandler"
)

type Provider struct {
	UserHandler *userhandler.UserHandler
}

func New(userHandler *userhandler.UserHandler) *Provider {
	return &Provider{
		UserHandler: userHandler,
	}
}
