package services_test

import (
	"errors"
	"net/http"
	"strconv"
	"test-sms-2-pro/config"
	"test-sms-2-pro/constant"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/internal/repositories/jsonFile"
	"test-sms-2-pro/internal/services"
	"test-sms-2-pro/loggers"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetPokemonByNameService(t *testing.T) {
	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		pokemonName   string
		mockData      map[string]interface{}
		expectSuccess models.PokemonResponse
		expectError   error
	}{
		{
			name:        "GetPokemonByNameServiceSuccess",
			pokemonName: "pokemon",
			mockData:    map[string]interface{}{"name": "pokemon"},

			expectSuccess: models.PokemonResponse{
				Status: http.StatusOK,
				Code:   strconv.Itoa(http.StatusOK),
				Data:   &map[string]interface{}{"name": "pokemon"},
			},
			expectError: nil,
		},
		{
			name:        "GetPokemonByNameServiceErrorNotFound",
			pokemonName: "pokemon2",
			mockData:    map[string]interface{}{"name": "pokemon"},

			expectSuccess: models.PokemonResponse{
				Status:  http.StatusOK,
				Code:    strconv.Itoa(http.StatusOK),
				Message: constant.UserRegisterSuccessMessage,
				Data:    &map[string]interface{}{"name": "pokemon"},
			},
			expectError: errors.New(constant.PokemonErrorMessageNotFound),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			pokemonRepo := jsonFile.NewPokemonRepositoryMock()
			switch tC.name {
			case "GetPokemonByNameServiceErrorNotFound":
				pokemonRepo.On("GetPokemonByName").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			default:
				pokemonRepo.On("GetPokemonByName").Return(tC.mockData, nil)
				break
			}
			userSvc := services.NewPokemonService(pokemonRepo)
			resp, err := userSvc.GetPokemonByNameService(tC.pokemonName)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestGetPokemonAbilityByNameService(t *testing.T) {
	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		pokemonName   string
		mockData      map[string]interface{}
		expectSuccess models.PokemonResponse
		expectError   error
	}{
		{
			name:        "GetPokemonAbilityByNameServiceSuccess",
			pokemonName: "pokemon",
			mockData:    map[string]interface{}{"abilities": "pokemon"},

			expectSuccess: models.PokemonResponse{
				Status: http.StatusOK,
				Code:   strconv.Itoa(http.StatusOK),
				Data:   &map[string]interface{}{"abilities": "pokemon"},
			},
			expectError: nil,
		},
		{
			name:        "GetPokemonAbilityByNameServiceErrorNotFoundPokemon",
			pokemonName: "pokemon2",
			mockData:    map[string]interface{}{"abilities": "pokemon"},

			expectSuccess: models.PokemonResponse{
				Status:  http.StatusOK,
				Code:    strconv.Itoa(http.StatusOK),
				Message: constant.PokemonErrorMessageNotFound,
				Data:    &map[string]interface{}{"abilities": "pokemon"},
			},
			expectError: errors.New(constant.PokemonErrorMessageNotFound),
		},
		{
			name:        "GetPokemonAbilityByNameServiceErrorNotFoundAbility",
			pokemonName: "pokemon2",
			mockData:    map[string]interface{}{"abilities": nil},

			expectSuccess: models.PokemonResponse{
				Status: http.StatusOK,
				Code:   strconv.Itoa(http.StatusOK),
				Data:   &map[string]interface{}{"abilities": "pokemon"},
			},
			expectError: errors.New(constant.PokemonAbilityErrorMessageNotFound),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			pokemonRepo := jsonFile.NewPokemonRepositoryMock()
			switch tC.name {
			case "GetPokemonAbilityByNameServiceErrorNotFoundPokemon":
				pokemonRepo.On("GetPokemonByName").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			case "GetPokemonAbilityByNameServiceErrorNotFoundAbility":
				pokemonRepo.On("GetPokemonByName").Return(tC.mockData, nil)
				break
			default:
				pokemonRepo.On("GetPokemonByName").Return(tC.mockData, nil)
				break
			}
			userSvc := services.NewPokemonService(pokemonRepo)
			resp, err := userSvc.GetPokemonAbilityByNameService(tC.pokemonName)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
