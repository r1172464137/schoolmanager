package model

import "gorm.io/gorm"

type Article struct {
	//ID        uint   `gorm:"type:int;primaryKey;not null"`
	gorm.Model
	Tile      string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Publisher uint   `gorm:"not null"`
}
