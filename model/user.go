package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	//ID             uint   `gorm:"type:int;primaryKey"`
	gorm.Model
	Uid            uint   `gorm:"unique"`
	Name           string `gorm:"size:5"`
	Username       string `gorm:"size:15;unique"`
	PasswordDigest string
	Capacity       bool   `gorm:"default:false"` //0学生 1老师
	College        string `gorm:"size:15"`
}

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func (studentUser *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	studentUser.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (studentUser *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(studentUser.PasswordDigest), []byte(password))
	return err == nil
}
