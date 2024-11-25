package jsonFile

import (
	"errors"
)

var Pokemons map[string]interface{}

type pokemonRepository struct {
}

// GetPokemonByName implements PokemonRepository.
func (p pokemonRepository) GetPokemonByName(pokemonName string) (map[string]interface{}, error) {

	if Pokemons[pokemonName] != nil {
		return Pokemons[pokemonName].(map[string]interface{}), nil
	}
	return nil, errors.New("pokemon not found")
}

func LoadPokemonsData(data map[string]interface{}) {
	Pokemons = data
}
func NewPokemonsRepository() PokemonRepository {
	return pokemonRepository{}
}
