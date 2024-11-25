package handlers

import (
	"net/http"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/internal/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	usersService services.UsersService
}

// LoginHandler implements UsersHandler.
func (u usersHandler) LoginHandler(c echo.Context) error {
	usersReq := new(models.UsersRequest)
	if err := c.Bind(usersReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(usersReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userResp, err := u.usersService.LoginService(*usersReq)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, userResp, "")
}

// RegisterHandler implements UsersHandler.
func (u usersHandler) RegisterHandler(c echo.Context) error {
	usersReq := new(models.UsersRequest)
	if err := c.Bind(usersReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(usersReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userResp, err := u.usersService.RegisterService(*usersReq)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusCreated, userResp, "")
}

func NewUsersHandler(usersService services.UsersService) UsersHandler {
	return usersHandler{usersService: usersService}
}
