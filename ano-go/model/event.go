package model

import (
	"gorm.io/gorm"
)

type Event struct {
	ID			uint			`jsom:"id" gorm:"primaryKey"`
	Name		string			`json:"name"`
	Description	string			`json:"description"`
	ImagePath	string			`json:"imagePath"`
	CreatedAt	uint64			`json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt	uint64			`json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt	gorm.DeletedAt	`json:"-" gorm:"index"`
}