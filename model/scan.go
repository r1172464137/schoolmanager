package model

import "gorm.io/gorm"

type Scan struct {
	gorm.Model
	//Id   uint   `gorm:"primaryKey"`
	Base string `gorm:"size:50"`
	Hash string `gorm:"size:50"`
	Last string `gorm:"size:50"`
}
