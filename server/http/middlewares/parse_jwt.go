package middlewares

import (
	"github.com/godvlpr/template"
	"github.com/godvlpr/template/errs"
	"github.com/godvlpr/template/server/http/custctx"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

var (
	userID             = "user_id"
	currentAccountType = "current_account_type"
)

func ParseToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _, err := getShopperIDFromJWT(c.Request())
		if err != nil {
			return c.JSON(errs.UnauthorizedErr.Code, errs.UnauthorizedErr)
		}

		accType, _, err := getCurrentAccTypeFromJWT(c.Request())
		if err != nil {
			return c.JSON(errs.UnauthorizedErr.Code, errs.UnauthorizedErr)
		}

		cc := &custctx.CustomContext{
			Context:          c,
			UserID:           userID,
			RequestOwnerType: accType,
		}

		return next(cc)
	}
}

func getCurrentAccTypeFromJWT(r *http.Request) (string, string, error) {
	return getFromJWT(r, currentAccountType)
}

func getShopperIDFromJWT(r *http.Request) (string, string, error) {
	return getFromJWT(r, userID)
}

func getFromJWT(r *http.Request, fieldType string) (string, string, error) {
	var tokenRaw string
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		tokenRaw = bearer[7:]
	}

	token, err := jwt.Parse(tokenRaw, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.GetConfig().GetAuthCopy().VerifyKey), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("cannot cast token.Claims to jwt.MapClaims")
	}

	var fieldRaw interface{}

	fieldRaw, ok = claims[fieldType]
	if !ok {
		return "", "", errors.New("shopper_id is absent in the jwt")
	}

	fieldValue, ok := fieldRaw.(string)
	if !ok {
		return "", "", errors.New("failed to cast shopper_id into string")
	}

	return fieldValue, token.Raw, nil
}
