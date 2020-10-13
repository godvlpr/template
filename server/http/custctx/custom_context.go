package custctx

import "github.com/labstack/echo"

//CustomContext - implementation of echo context and custom fields which will be parsed in middleware
type CustomContext struct {
	echo.Context
	UserID           string
	RequestOwnerType string
}
