package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"microservice/pkg/utils"
	"net/http"
)

type FastPayForm struct {
	UserId   uint    `json:"UserId" validate:"required,gt=0"`
	Amount   float64 `json:"Amount" validate:"required,gt=0"`
	Provider string  `json:"Provider" validate:"required,provider"`
	Currency string  `json:"Currency" validate:"required,currency"`
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

	c.JSON(http.StatusOK, gin.H{"message": "FastPay successful", "data": form})
	return
}
