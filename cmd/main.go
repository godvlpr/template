package main

import (
	app "github.com/godvlpr/template"
	"github.com/godvlpr/template/docs"
)

// @title Template
// @version 1.0
// @securityDefinitions.apiKey bearerToken
// @in Header
// @name Authorization
func main() {
	docs.SwaggerInfo.Host = "localhost:1123"
	app.App()
}
