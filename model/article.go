package model

import "gorm.io/gorm"

type Article struct {
	//ID        uint   `gorm:"type:int;primaryKey;not null"`
	gorm.Model
	Tile      string `gorm:"type:varchar(20);not null"`
	Content   string `gorm:"type:varchar(50);not null"`
	Publisher uint   `gorm:"type:int;not null"`
}
