package handlers

import (
	"net/http"
	"test-sms-2-pro/errs"

	"github.com/labstack/echo/v4"
)

type UsersHandler interface {
	LoginHandler(c echo.Context) error
	RegisterHandler(c echo.Context) error
}
type PokemonHandler interface {
	SearchPokemonByNameHandler(c echo.Context) error
	SearchAbilityPokemonByNameHandler(c echo.Context) error
}

func HandlerError(err error) *echo.HTTPError {
	switch e := err.(type) {
	case errs.AppError:
		return echo.NewHTTPError(e.Code, e.Message)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
