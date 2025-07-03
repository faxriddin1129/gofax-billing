package models

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	UserID int64     `json:"user_id" gorm:"column:user_id"`
	Token  string    `json:"token" gorm:"type:varchar(255)"`
	Expire time.Time `json:"expire"`
	Ip     string    `json:"ip" gorm:"type:varchar(255)"`
	Device string    `json:"device"`
}

func (Token) TableName() string {
	return "tokens"
}

func (t *Token) CreateAccessToken(db *gorm.DB) error {
	return db.Create(t).Error
}

func TokenExists(token string) (uint, bool) {
	var result struct {
		UserID uint
	}

	err := DB.Where("token = ? AND expire > ?", token, time.Now()).
		Select("user_id").
		Scan(&result).Error

	if err != nil || result.UserID == 0 {
		return 0, false
	}
	return result.UserID, true
}
