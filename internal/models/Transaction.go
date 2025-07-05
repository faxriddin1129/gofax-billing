package models

import (
	"gorm.io/gorm"
	"microservice/pkg/utils"
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
	OrderId       string  `json:"OrderId" gorm:"type:string"`
	ProductId     string  `json:"ProductId" gorm:"type:string"`
	ReturnUrl     string  `json:"ReturnUrl" gorm:"type:string"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func TransactionGetAll() []Transaction {
	var transactions []Transaction
	utils.DB.Find(&transactions)
	return transactions
}
