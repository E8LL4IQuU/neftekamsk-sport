package model

import (
	"gorm.io/gorm"
)

type Event struct {
	ID			uint
	Title		string
	Description	string
	ImagePath	string
	Date		string
	CreatedAt	uint64
	UpdatedAt	uint64
	DeletedAt	gorm.DeletedAt
}
