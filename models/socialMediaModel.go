package models

type SocialMedia struct{
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
}