package handlers

import (
	"net/http"
	"test-sms-2-pro/internal/services"

	"github.com/labstack/echo/v4"
)

type pokemonHandler struct {
	pokemonService services.PokemonServices
}

// SearchAbilityPokemonByNameHandler implements PokemonHandler.
func (p pokemonHandler) SearchAbilityPokemonByNameHandler(c echo.Context) error {
	name := c.Param("name")
	pokemonresp, err := p.pokemonService.GetPokemonAbilityByNameService(name)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, pokemonresp, "")
}

// SearchPokemonByNameHandler implements PokemonHandler.
func (p pokemonHandler) SearchPokemonByNameHandler(c echo.Context) error {
	name := c.Param("name")
	pokemonresp, err := p.pokemonService.GetPokemonByNameService(name)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, pokemonresp, "")
}

func NewPokemonsHandler(pokemonService services.PokemonServices) PokemonHandler {
	return pokemonHandler{pokemonService: pokemonService}
}
