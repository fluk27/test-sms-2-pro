package jsonFile

type PokemonRepository interface {
	GetPokemonByName(pokemonName string) (map[string]interface{}, error)
}
