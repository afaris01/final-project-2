package models

type Comment struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	UserID  uint
	User    *User
	PhotoID uint `json:"photo_id" form:"photo_id"`
	Photo   *Photo
	Content string `gorm:"not null" json"content" form:"content" valid:"required~Comment is required"`
	TimeModel
}
