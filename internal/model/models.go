package model

import (
	"gorm.io/gorm"
	"time"
)

// Model same as in gorm but add json tags
type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	gorm.Model
	Account string
	Password string
	Type string
}

type Good struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Description string
	Count uint
}

type Cart struct {
	gorm.Model
	User User `gorm:"foreignKey:ID"`
}

type CartGood struct {
	Model
	Cart  Cart `gorm:"foreignKey:ID;"` // constraint:OnDelete:CASCADE;
	Good  Good `gorm:"foreignKey:ID;"`
	Count uint
}
