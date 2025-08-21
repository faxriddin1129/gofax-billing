package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gofax-billing/internal/constants"
	"gofax-billing/internal/models"
	"gofax-billing/pkg/click"
	"gofax-billing/pkg/octo"
	"gofax-billing/pkg/payme"
	"gofax-billing/pkg/utils"
	"net/http"
	"time"
)

type FastPayForm struct {
	UserId    uint    `json:"UserId" validate:"required,gt=0"`
	Amount    float64 `json:"Amount" validate:"required,gt=0"`
	Provider  string  `json:"Provider" validate:"required,provider"`
	Currency  string  `json:"Currency" validate:"required,currency"`
	OrderId   string  `json:"OrderId" validate:"required,gt=0"`
	Platform  string  `json:"Platform" validate:"required,gt=0"`
	ProductId string  `json:"ProductId" validate:"required,gt=0"`
	ReturnUrl string  `json:"ReturnUrl" validate:"url"`
	Email     string  `json:"Email" validate:"required,email"`
	Phone     string  `json:"Phone"`
}

func FastPayValidate(c *gin.Context) {
	var form FastPayForm

	if err := c.ShouldBindJSON(&form); err != nil {
		utils.RespondJson(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	err := validate.Struct(form)
	msg := ""
	if err != nil {
		errorMessage := map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s %s", err.Field(), err.Tag())
			errorMessage[err.Field()] = errMsg
			if msg == "" {
				msg = errMsg
			}
		}
		utils.RespondJson(c, errorMessage, http.StatusBadRequest, msg)
		return
	}

	finalAmount := form.Amount
	if form.Provider == constants.ProviderPayme {
		finalAmount = form.Amount * 100
	}

	transaction := models.Transaction{
		Type:        constants.TYPE_FAST_PAY,
		Status:      constants.STATUS_PENDING,
		Currency:    form.Currency,
		Provider:    form.Provider,
		Amount:      finalAmount,
		State:       0,
		Reason:      0,
		UUID:        uuid.New().String(),
		CreateTime:  time.Now().Unix(),
		PerformTime: 0,
		OrderId:     form.OrderId,
		ReturnUrl:   form.ReturnUrl,
		ProductId:   form.ProductId,
		Email:       form.Email,
		Phone:       form.Phone,
		UserId:      form.UserId,
		Platform:    form.Platform,
	}

	err = utils.DB.Create(&transaction).Error
	if err != nil {
		utils.RespondJson(c, nil, http.StatusInternalServerError, "Internal server error. Transaction failed save")
		return
	}

	if transaction.Provider == constants.ProviderClick {
		data, code, msg := click.GenerateShopApiLink(&transaction)
		utils.RespondJson(c, data, code, msg)
		return
	}

	if transaction.Provider == constants.ProviderPayme {
		data, code, msg := payme.GenerateShopApiLink(&transaction)
		utils.RespondJson(c, data, code, msg)
		return
	}

	if transaction.Provider == constants.ProviderOcto {
		data, code, msg := octo.GenerateShopApiLink(&transaction)
		utils.RespondJson(c, data, code, msg)
		return
	}

	utils.RespondJson(c, nil, http.StatusNotFound, "Provider not found or inactive")
	return
}
