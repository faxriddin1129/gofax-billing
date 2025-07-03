package models

import (
	"gorm.io/gorm"
	"microservice/pkg/utils"
	"time"
)

type Token struct {
	gorm.Model
	UserID uint      `json:"UserId" gorm:"column:user_id"`
	Token  string    `json:"Token" gorm:"type:varchar(255)"`
	Expire time.Time `json:"Expire"`
	Ip     string    `json:"Ip" gorm:"type:varchar(255)"`
	Device string    `json:"Device"`
}

func (Token) TableName() string {
	return "tokens"
}

func TokenExists(token string) (uint, bool) {
	var result struct {
		UserID uint
	}

	err := utils.DB.Where("token = ? AND expire > ?", token, time.Now()).
		Select("user_id").
		Scan(&result).Error

	if err != nil || result.UserID == 0 {
		return 0, false
	}
	return result.UserID, true
}
