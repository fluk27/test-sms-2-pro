package services_test

import (
	"errors"
	"net/http"
	"strconv"
	"test-sms-2-pro/config"
	"test-sms-2-pro/constant"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/internal/repositories/db"
	"test-sms-2-pro/internal/services"
	"test-sms-2-pro/loggers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLoginUser(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		request       models.UsersRequest
		mockData      models.UsersRepository
		expectSuccess models.UsersResponse
		expectError   error
	}{
		{
			name: "TestLoginUserSuccess",
			request: models.UsersRequest{
				Username: "ariyavas",
				Password: "ariyavas",
			},
			mockData: models.UsersRepository{
				ID:               1,
				Username:         "ariyavas",
				Password:         "$2a$10$p6ONl3UzmRzlTMPO0KM4UOCLOG49ToYG6xmj3G.8SV1y9JsFm1oNq",
				Email:            nil,
				IsActive:         true,
				CreatedTimestamp: time.Now(),
				UpdatedTimestamp: time.Now(),
			},

			expectSuccess: models.UsersResponse{
				Message: constant.UserLoginSuccessMessage,
				Data: &models.UsersData{
					AccessToken:  "",
					RefreshToken: "",
				},
			},
			expectError: nil,
		},
		{
			name: "TestLoginUserErrorInternalServerError",
			request: models.UsersRequest{
				Username: "ariyavas",
				Password: "ariyavas",
			},
			mockData: models.UsersRepository{
				ID:               1,
				Username:         "ariyavas",
				Password:         "ariyavas",
				Email:            nil,
				IsActive:         true,
				CreatedTimestamp: time.Now(),
				UpdatedTimestamp: time.Now(),
			},

			expectSuccess: models.UsersResponse{
				Message: constant.UserLoginSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.MessageInternalServerError),
		},
		{
			name: "TestLoginUserFindNotFound",
			request: models.UsersRequest{
				Username: "ariyavas2",
				Password: "ariyavas2",
			},
			mockData: models.UsersRepository{},

			expectSuccess: models.UsersResponse{
				Message: constant.UserLoginSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.UserErrorMessageNotFound),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			userRepo := db.NewUsersRepositoryMock()

			switch tC.name {
			case "TestLoginUserFindNotFound":
				userRepo.On("GetUserByUserName").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			case "TestLoginUserErrorInternalServerError":
				userRepo.On("GetUserByUserName").Return(tC.mockData, errors.New("error mock"))
				break
			default:
				userRepo.On("GetUserByUserName").Return(tC.mockData, nil)
				break
			}

			userSvc := services.NewUsersService(userRepo)
			resp, err := userSvc.LoginService(tC.request)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {
				assert.NotEmpty(t, resp.Data.AccessToken)
				//assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestRegisterUser(t *testing.T) {
	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		request       models.UsersRequest
		mockData      models.UsersRepository
		expectSuccess models.UsersResponse
		expectError   error
	}{
		{
			name: "RegisterUserSuccess",
			request: models.UsersRequest{
				Username: "ariyavas",
				Password: "ariyavas",
			},
			mockData: models.UsersRepository{},

			expectSuccess: models.UsersResponse{
				Status:  http.StatusCreated,
				Code:    strconv.Itoa(http.StatusCreated),
				Message: constant.UserRegisterSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name: "RegisterUserErrorDuplicate",
			request: models.UsersRequest{
				Username: "ariyavas",
				Password: "ariyavas",
			},
			mockData: models.UsersRepository{
				ID:               1,
				Username:         "ariyavas",
				Password:         "ariyavas",
				Email:            nil,
				IsActive:         true,
				CreatedTimestamp: time.Now(),
				UpdatedTimestamp: time.Now(),
			},

			expectSuccess: models.UsersResponse{
				Message: constant.UserRegisterSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.UserErrorsMessageConflict),
		},
		{
			name: "RegisterUserErrorInternalServerError",
			request: models.UsersRequest{
				Username: "ariyavas",
				Password: "ariyavas",
			},
			mockData: models.UsersRepository{
				ID:               1,
				Username:         "ariyavas",
				Password:         "ariyavas",
				Email:            nil,
				IsActive:         true,
				CreatedTimestamp: time.Now(),
				UpdatedTimestamp: time.Now(),
			},

			expectSuccess: models.UsersResponse{
				Message: constant.UserRegisterSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.MessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			userRepo := db.NewUsersRepositoryMock()
			switch tC.name {
			case "RegisterUserErrorInternalServerError":
				userRepo.On("CreateUser").Return(errors.New("gorm error"))
				break
			case "RegisterUserErrorDuplicate":
				userRepo.On("CreateUser").Return(gorm.ErrDuplicatedKey)
				break
			default:
				userRepo.On("CreateUser").Return(nil)
				break
			}
			userSvc := services.NewUsersService(userRepo)
			resp, err := userSvc.RegisterService(tC.request)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
