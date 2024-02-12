package model

import (
	"gorm.io/gorm"
)

type News struct {
	ID	uint64
	Title string
	Description string
	ImagePath string
	CreatedAt uint64
	UpdatedAt uint64
	DeletedAt gorm.DeletedAt
}