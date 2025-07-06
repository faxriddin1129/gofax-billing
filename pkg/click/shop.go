package click

import (
	"microservice/internal/models"
	"strconv"
)

func GenerateShopApiLink(transaction *models.Transaction) interface{} {
	amount := strconv.FormatFloat(transaction.Amount, 'f', 2, 64)
	merchantId := MERCHANT_ID
	merchantUserId := MERCHANT_USER_ID
	serviceId := SERVICE_ID
	merchantTransId := strconv.Itoa(int(transaction.ID))
	returnUrl := transaction.ReturnUrl
	serviceUrl := SERVICE_URL

	link := serviceUrl + "?amount=" + amount + "&merchant_id=" + merchantId + "&merchant_user_id=" + merchantUserId + "&service_id=" + serviceId + "&merchant_trans_id=" + merchantTransId + "&return_url" + returnUrl

	return map[string]interface{}{
		"ID":     transaction.ID,
		"Link":   link,
		"Method": "POST",
	}
}
