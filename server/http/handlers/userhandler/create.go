package userhandler

import (
	"github.com/godvlpr/template/errs"
	"github.com/godvlpr/template/models"
	"github.com/godvlpr/template/server/datamodels"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ShowAccount godoc
// @Security bearerAuth
// @Summary Show an account
// @Tags users
// @Consume application/json
// @Param JSON body datamodels.User true "it is id"
// @Description get test user model
// @Accept  json
// @Produce  json
// @Success 200 {object} datamodels.User
// @Header 200 {string} Token "JWT token"
// @Failure 500 {object} errs.ErrWrapper
// @Router /users [post]
func (h UserHandler) Create(c echo.Context) error {
	var req datamodels.User

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("Filed to unmarshal request body")
		return c.JSON(http.StatusBadRequest, errs.IncorrectBodyErr)
	}

	user := models.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}

	user, err = h.userService.Create(user)
	if err != nil {
		h.log.WithError(err).Error("Filed to create user")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	resp := datamodels.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
	}

	return c.JSON(http.StatusOK, resp)
}
