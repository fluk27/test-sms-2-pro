package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/loggers"
	"test-sms-2-pro/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func JWTCustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		req := c.Request().Header
		auth := req.Get("Authorization")
		jwtToken := strings.Split(auth, " ")
		if len(jwtToken) < 2 {
			msgError := errors.New("jwtToken is empty.")
			Unauthorized := models.UsersResponse{
				Status:  http.StatusUnauthorized,
				Code:    strconv.Itoa(http.StatusUnauthorized),
				Message: msgError.Error(),
			}
			return c.JSONPretty(http.StatusUnauthorized, Unauthorized, "")
		}
		_, err := utils.DecodeJWT(jwtToken[1])
		if err != nil {
			loggers.Error("decode jWT error", zap.Error(err))
			Unauthorized := models.UsersResponse{
				Status:  http.StatusUnauthorized,
				Code:    strconv.Itoa(http.StatusUnauthorized),
				Message: strings.Split(err.Error(), ":")[1],
			}
			return c.JSONPretty(http.StatusUnauthorized, Unauthorized, "")
		}
		return next(c)
	}
}
