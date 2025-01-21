package repositories

import "gorm.io/gorm"

type User interface {
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}
