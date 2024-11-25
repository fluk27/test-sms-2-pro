package models

import "time"

type UsersRepository struct {
	ID               int    `gorm:"primaryKey;autoIncrement"`
	Username         string `gorm:"unique;not null"`
	Password         string `gorm:"not null"`
	Email            *string
	IsActive         bool      `gorm:"default:true"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime"`
	UpdatedTimestamp time.Time `gorm:"autoUpdateTime"`
}

func (UsersRepository) TableName() string {
	return "users"
}
