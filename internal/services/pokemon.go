package services

import (
	"fmt"
	"net/http"
	"strconv"
	"test-sms-2-pro/constant"
	"test-sms-2-pro/errs"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/internal/repositories/jsonFile"
	"test-sms-2-pro/loggers"

	"go.uber.org/zap"
)

type pokemonService struct {
	pokemonrepo jsonFile.PokemonRepository
}

// GetPokemonAbilityByNameService implements PokemonServices.
func (p pokemonService) GetPokemonAbilityByNameService(name string) (models.PokemonResponse, error) {
	pokemondata, err := p.pokemonrepo.GetPokemonByName(name)
	if err != nil {
		loggers.Error(fmt.Sprintf("GetPokemonByName=%v", err.Error()),
			zap.String("type", "repo"),
			zap.String("name", name),
			zap.Error(err))
		return models.PokemonResponse{}, errs.NewNotFoundError(constant.PokemonErrorMessageNotFound)

	}
	if pokemondata["abilities"] != nil {
		result := make(map[string]interface{})
		result["abilities"] = pokemondata["abilities"]
		return models.PokemonResponse{
			Status: http.StatusOK,
			Code:   strconv.Itoa(http.StatusOK),
			Data:   &result,
		}, nil
	} else {
		return models.PokemonResponse{}, errs.NewNotFoundError(constant.PokemonAbilityErrorMessageNotFound)
	}
}

// GetPokemonByNameService implements PokemonServices.
func (p pokemonService) GetPokemonByNameService(name string) (models.PokemonResponse, error) {
	pokemondata, err := p.pokemonrepo.GetPokemonByName(name)
	if err != nil {
		loggers.Error(fmt.Sprintf("GetPokemonByName=%v", err.Error()),
			zap.String("type", "repo"),
			zap.String("name", name),
			zap.Error(err))
		return models.PokemonResponse{}, errs.NewNotFoundError(constant.PokemonErrorMessageNotFound)

	}
	return models.PokemonResponse{
		Status: http.StatusOK,
		Code:   strconv.Itoa(http.StatusOK),
		Data:   &pokemondata,
	}, nil
}

func NewPokemonService(pokemonrepo jsonFile.PokemonRepository) PokemonServices {
	return pokemonService{pokemonrepo: pokemonrepo}
}
