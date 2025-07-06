package payme

import (
	"encoding/base64"
	"microservice/internal/models"
	"strconv"
)

func GenerateShopApiLink(transaction *models.Transaction) interface{} {
	merchantId := MERCHANT_ID
	serviceUrl := SERVICE_URL
	amount := strconv.FormatFloat(transaction.Amount, 'f', 2, 64)
	orderId := strconv.Itoa(int(transaction.ID))
	returnUrl := transaction.ReturnUrl
	language := "ru"

	paramStr := "m=" + merchantId + ";ac.order_id=" + orderId + ";a=" + amount + ";c=" + returnUrl + ";l=" + language
	encoded := base64.StdEncoding.EncodeToString([]byte(paramStr))
	link := serviceUrl + "/" + encoded

	return map[string]interface{}{
		"ID":     transaction.ID,
		"Link":   link,
		"Method": "GET",
	}
}
