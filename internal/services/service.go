package services

import (
	"test-sms-2-pro/internal/models"
)

type UsersService interface {
	LoginService(req models.UsersRequest) (models.UsersResponse, error)
	RegisterService(req models.UsersRequest) (models.UsersResponse, error)
}
type PokemonServices interface {
	GetPokemonByNameService(name string) (models.PokemonResponse, error)
	GetPokemonAbilityByNameService(name string) (models.PokemonResponse, error)
}