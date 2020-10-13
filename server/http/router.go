package http

import (
	"github.com/godvlpr/template/server/http/handlers"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Router(handlers *handlers.Provider) *echo.Echo {
	router := echo.New()
	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-auth", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposeHeaders:    []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	middleware.DefaultLoggerConfig.Format = `{"TIME":"${time_rfc3339_nano}","REMOTE_IP":"${remote_ip}",` +
		`"HOST":"${host}","METHOD":"${method}","URI":"${uri}","USER_AGENT":"${user_agent}",` +
		`"STATUS":${status},"ERROR":"${error}","LATENCY_HUMAN":"${latency_human}"`

	router.Use(
		//middlewares.ParseToken,
		cors,
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.DefaultLoggerConfig),
	)

	router.POST("/users", handlers.UserHandler.Create)
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, `{"status": "ok"}`)
	})

	return router
}
