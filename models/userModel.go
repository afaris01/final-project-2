package models

import (
	"final-project-2/helpers"
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)


type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is invalid"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length 6 characters"`
	Age      uint   `gorm:"not null" json:"age" form:"age" valid:"required~Age is required"`
	TimeModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age < 8 {
		err = errors.New("Minimum age to register is 8")
		return err
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}