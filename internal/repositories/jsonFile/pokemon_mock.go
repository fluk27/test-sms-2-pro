package jsonFile

import (
	"github.com/stretchr/testify/mock"
)

type mockPokemonRepository struct {
	mock.Mock
}

// GetPokemonByName implements PokemonRepository.
func (mockUsersRepo *mockPokemonRepository) GetPokemonByName(pokemonName string) (map[string]interface{}, error) {
	args := mockUsersRepo.Called()
	return args.Get(0).(map[string]interface{}), args.Error(1)
}
func NewPokemonRepositoryMock() *mockPokemonRepository {
	return &mockPokemonRepository{}
}
