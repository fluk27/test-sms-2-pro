package db

import (
	"test-sms-2-pro/internal/models"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

// CreateUser implements UsersRepository.
func (u usersRepository) CreateUser(req models.UsersRepository) error {
	if err := u.db.Create(&req).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByUserNameAndPassword implements UsersRepository.
func (u usersRepository) GetUserByUserName(username string) (models.UsersRepository, error) {
	userRepoResp := models.UsersRepository{}
	if err := u.db.Where("username", username).First(&userRepoResp).Error; err != nil {
		return userRepoResp, err
	}

	return userRepoResp, nil
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return usersRepository{db: db}
}
