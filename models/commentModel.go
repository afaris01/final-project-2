package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	UserID  uint
	User    *User
	PhotoID uint `json:"photo_id" form:"photo_id"`
	Photo   *Photo
	Content string `gorm:"not null" json:"content" form:"content" valid:"required~Comment is required"`
	TimeModel
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}