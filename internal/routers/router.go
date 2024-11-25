package routers

import (
	"net/http"
	"test-sms-2-pro/internal/handlers"
	"test-sms-2-pro/internal/services"
	customMid "test-sms-2-pro/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(usersSvc services.UsersService, pokemonSvc services.PokemonServices) *echo.Echo {
	e := echo.New()
	v1 := e.Group("/api/v1")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	pokemonRouter(v1, pokemonSvc)
	//users
	usersHandle := handlers.NewUsersHandler(usersSvc)
	api := v1.Group("/user")
	api.POST("/register", usersHandle.RegisterHandler)
	api.POST("/login", usersHandle.LoginHandler)
	return e
}
func heathCheck(c echo.Context) error {
	return c.String(http.StatusOK, "pokemon running.")
}
func pokemonRouter(eg *echo.Group, pokemonSvc services.PokemonServices) {
	pokemonHandle := handlers.NewPokemonsHandler(pokemonSvc)
	api := eg.Group("/pokemon")
	api.Use(customMid.JWTCustomMiddleware)
	api.GET("/:name", pokemonHandle.SearchPokemonByNameHandler)
	api.GET("/:name/ability", pokemonHandle.SearchAbilityPokemonByNameHandler)
}
