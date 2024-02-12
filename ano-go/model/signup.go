package model

import (
	"gorm.io/gorm"
)

type Signup struct {
	ID          uint64
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	CreatedAt   uint64
	DeletedAt   gorm.DeletedAt
}
