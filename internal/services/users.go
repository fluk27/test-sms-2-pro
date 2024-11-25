package services

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"test-sms-2-pro/constant"
	"test-sms-2-pro/errs"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/internal/repositories/db"
	"test-sms-2-pro/loggers"
	"test-sms-2-pro/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type usersService struct {
	usersRepo db.UsersRepository
}

var (
	accessTokenRedisKey  = "accessToken"
	refreshTokenRedisKey = "refreshToken"
	userDetailRedisKey   = "userDetail"
	accessTokenTime      = time.Minute * 60
	refreshTokenTime     = time.Hour * 24 * 90
)

// LoginService implements UsersService.
func (u usersService) LoginService(req models.UsersRequest) (models.UsersResponse, error) {
	usersData, err := u.usersRepo.GetUserByUserName(req.Username)
	if err != nil {
		loggers.Error(fmt.Sprintf("GetUserByUserName error=%v", err.Error()),
			zap.String("type", "repo"),
			zap.String("username", req.Username),
			zap.String("password", req.Password),
			zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return models.UsersResponse{}, errs.NewNotFoundError(constant.UserErrorMessageNotFound)
		} else {
			return models.UsersResponse{}, errs.NewInternalServerError(constant.MessageInternalServerError)
		}
	}
	if !utils.CheckPasswordHash(req.Password, usersData.Password) {
		return models.UsersResponse{}, errs.NewBadRequest(constant.UserErrorInvalidUsernameOrPassword)
	}
	accessToken, err := utils.CreateJWT(strconv.Itoa(usersData.ID), usersData.Username, usersData.IsActive, accessTokenTime)
	if err != nil {
		loggers.Error(fmt.Sprintf("cerate access token error=%s", err.Error()),
			zap.String("type", "utils"),
			zap.Error(err),
		)
		return models.UsersResponse{}, errs.NewInternalServerError(constant.MessageInternalServerError)
	}
	refreshToken, err := utils.CreateJWT(strconv.Itoa(usersData.ID), usersData.Username, usersData.IsActive, refreshTokenTime)
	if err != nil {
		loggers.Error(fmt.Sprintf("cerate refresh token error=%s", err.Error()),
			zap.String("type", "utils"),
			zap.Error(err),
		)
		return models.UsersResponse{}, errs.NewInternalServerError(constant.MessageInternalServerError)
	}
	userData := models.UsersData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return models.UsersResponse{
		Status: http.StatusOK,
		Code:   strconv.Itoa(http.StatusOK),
		Data:   &userData,
	}, nil
}

// RegisterService implements UsersService.
func (u usersService) RegisterService(req models.UsersRequest) (models.UsersResponse, error) {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		loggers.Error(fmt.Sprintf("HashPassword=%v", err.Error()),
			zap.String("type", "utils"),
			zap.String("passwordReq", req.Password),
			zap.Error(err))
		return models.UsersResponse{}, errs.NewInternalServerError(constant.MessageInternalServerError)
	}
	usersRepReq := models.UsersRepository{
		Username: req.Username,
		Password: hashPassword,
	}
	err = u.usersRepo.CreateUser(usersRepReq)
	if err != nil {
		loggers.Error(fmt.Sprintf("CreateUser=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Any("usersRepReq", usersRepReq),
			zap.Error(err))
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return models.UsersResponse{}, errs.NewConflict(constant.UserErrorsMessageConflict)

		} else {

			return models.UsersResponse{}, errs.NewInternalServerError(constant.MessageInternalServerError)
		}

	}
	return models.UsersResponse{
		Status:  http.StatusCreated,
		Code:    strconv.Itoa(http.StatusCreated),
		Message: constant.UserRegisterSuccessMessage,
	}, nil
}

func NewUsersService(usersRepo db.UsersRepository) UsersService {
	return usersService{usersRepo: usersRepo}
}
