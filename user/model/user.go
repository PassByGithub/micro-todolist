package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWord string
}

const PassWordCost = 12 //密码加密难度

//Register: Password encrypted
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
}

//Login: Password compare
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return err == nil
}
