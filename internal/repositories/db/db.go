package db

import (
	"test-sms-2-pro/internal/models"
)

type UsersRepository interface {
	GetUserByUserName(username string) (models.UsersRepository, error)
	CreateUser(req models.UsersRepository) error
}
