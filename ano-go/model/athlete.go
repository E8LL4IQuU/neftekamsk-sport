package model

import (
	"gorm.io/gorm"
)

type Athlete struct {
	ID          uint64
	Name        string
	Description string
	ImagePath   string
	CreatedAt   uint64
	UpdatedAt   uint64
	DeletedAt   gorm.DeletedAt
}
