package userhandler

import (
	"github.com/godvlpr/template/config"
	"github.com/godvlpr/template/service/usersservice"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService usersservice.UserService
	log         *logrus.Entry
}

func NewUserHandler(userService usersservice.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         config.GetConfig().GetLogCopy().Entry,
	}
}
