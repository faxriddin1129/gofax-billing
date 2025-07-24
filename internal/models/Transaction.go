package models

import (
	"gofax-billing/pkg/utils"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Type          string  `json:"Type" gorm:"type:varchar(32)"`
	Status        int8    `json:"Status" gorm:"type:int"`
	Provider      string  `json:"Provider" gorm:"type:varchar(32)"`
	Currency      string  `json:"Currency" gorm:"type:varchar(3)"`
	Amount        float64 `json:"Amount" gorm:"type:float"`
	PaymentStatus int8    `json:"PaymentStatus" gorm:"type:int"`
	State         int8    `json:"State" gorm:"type:int"`
	CreateTime    int64   `json:"CreateTime" gorm:"type:int"`
	PerformTime   int64   `json:"PerformTime" gorm:"type:int"`
	CancelTime    int64   `json:"CancelTime" gorm:"type:int"`
	TransactionId uint    `json:"TransactionId" gorm:"type:int"`
	Reason        int     `json:"Reason" gorm:"type:int"`
	UUID          string  `json:"Uuid" gorm:"type:string"`
	OrderId       string  `json:"OrderId" gorm:"type:varchar(64)"`
	ProductId     string  `json:"ProductId" gorm:"type:varchar(64)"`
	ReturnUrl     string  `json:"ReturnUrl" gorm:"type:varchar(64)"`
	Phone         string  `json:"Phone" gorm:"type:varchar(32)"`
	Email         string  `json:"Email" gorm:"type:varchar(64)"`
	UserId        uint    `json:"UserId"`
	CardNumber    string  `json:"CardNumber"`
	CardExpire    string  `json:"CardExpire"`
	CardType      string  `json:"CardType"`
	CardCvv       string  `json:"CardCvv"`
	Platform      string  `json:"Platform"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func TransactionGetById(Id int64) Transaction {
	var transaction Transaction
	utils.DB.First(&transaction, Id)
	return transaction
}

func TransactionGetByUUID(Id string) Transaction {
	var transaction Transaction
	utils.DB.Where("UUID = ?", Id).First(&transaction)
	return transaction
}

func TransactionUpdate(t *Transaction) (int64, error) {
	res := utils.DB.Where("ID=?", t.ID).Updates(t)
	return res.RowsAffected, res.Error
}
