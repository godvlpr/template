package app

import (
	"github.com/godvlpr/template/config"
	"github.com/godvlpr/template/server/http"
	"github.com/godvlpr/template/server/http/handlers"
	"github.com/godvlpr/template/server/http/handlers/userhandler"
	"github.com/godvlpr/template/service/usersservice"
	"github.com/godvlpr/template/storage"
)

func App() {
	storage := storage.New()

	userService := usersservice.New(storage)

	userHandler := userhandler.NewUserHandler(userService)

	handlerProvider := handlers.New(userHandler)

	router := http.Router(handlerProvider)

	config.GetConfig().GetLogCopy().Fatal(http.NewServer().Listen(router))
}
