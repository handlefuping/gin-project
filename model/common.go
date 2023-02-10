package model

import (
	"gorm.io/gorm"
	"time"
)

var DBInstance *gorm.DB

type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}