package model

import "gorm.io/gorm"

//	type Dorm struct {
//		ID       uint       `gorm:"type:int;primaryKey;not null"`
//		Date     *time.Time `gorm:"not null"`
//		Student  uint       `gorm:"type:int;not null"`
//		ImageSrc string     `gorm:"type:varchar(50),not null"`
//	}
type Dorm struct {
	//ID       uint
	gorm.Model
	//Date     *time.Time `gorm:"not null"`
	Student  uint   `gorm:"not null"`
	ImageSrc string `gorm:"not null"`
}
