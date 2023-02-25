package model

import (
	"gorm.io/gorm"
)

type Leave struct {
	//ID     uint      `gorm:"type:int;primaryKey;not null"`
	gorm.Model
	Uid    uint   `gorm:"column:uid;not null"`
	Name   string `gorm:"size:5"`
	Time   uint   `gorm:"size:7"`
	Reason string `gorm:"type:varchar(50)"`
	Status uint   `gorm:"type:bool;default:2"` //0 申请成功 1申请失败  2 申请中
}
