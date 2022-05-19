package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct{
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserID   uint
	User     *User
	TimeModel
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}