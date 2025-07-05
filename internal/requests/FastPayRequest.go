package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"microservice/internal/constants"
	"microservice/internal/models"
	"microservice/pkg/utils"
	"net/http"
	"time"
)

type FastPayForm struct {
	UserId   uint    `json:"UserId" validate:"required,gt=0"`
	Amount   float64 `json:"Amount" validate:"required,gt=0"`
	Provider string  `json:"Provider" validate:"required,provider"`
	Currency string  `json:"Currency" validate:"required,currency"`
	OrderId  string  `json:"OrderId" validate:"required,gt=0"`
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

	transaction := models.Transaction{
		Type:        constants.TYPE_FAST_PAY,
		Status:      constants.STATUS_PENDING,
		Currency:    form.Currency,
		Provider:    form.Provider,
		Amount:      form.Amount * 100,
		State:       constants.STATE_PENDING,
		Reason:      constants.REASON_PENDING,
		UUID:        uuid.New().String(),
		CreateTime:  time.Now().Unix(),
		PerformTime: time.Now().Unix(),
		OrderId:     form.OrderId,
	}

	err = utils.DB.Create(&transaction).Error
	if err != nil {
		utils.RespondJson(c, nil, http.StatusInternalServerError, "Internal server error. Transaction failed save")
		return
	}

	utils.RespondJson(c, transaction, http.StatusOK, "Transaction created successfully")
	return
}
