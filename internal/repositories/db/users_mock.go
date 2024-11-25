package db

import (
	"test-sms-2-pro/internal/models"

	"github.com/stretchr/testify/mock"
)

type mockUsersRepository struct {
	mock.Mock
}

// CreateUser implements UserRepository.
func (mockUsersRepo *mockUsersRepository) CreateUser(req models.UsersRepository) error {
	args := mockUsersRepo.Called()
	return args.Error(0)
}

// GetUserByUserName implements UserRepository.
func (mockUsersRepo *mockUsersRepository) GetUserByUserName(username string) (models.UsersRepository, error) {
	args := mockUsersRepo.Called()
	return args.Get(0).(models.UsersRepository), args.Error(1)
}
func NewUsersRepositoryMock() *mockUsersRepository {
	return &mockUsersRepository{}
}
