package click

import (
	"microservice/internal/models"
	"strconv"
)

func GenerateShopApiLink(transaction *models.Transaction) interface{} {
	amount := strconv.FormatFloat(transaction.Amount, 'f', 2, 64)
	merchantId := MerchantID
	merchantUserId := MerchantUserID
	serviceId := ServiceID
	merchantTransId := strconv.Itoa(int(transaction.ID))
	returnUrl := transaction.ReturnUrl
	serviceUrl := ServiceURL
	logo := Logo
	link := serviceUrl + "?amount=" + amount + "&merchant_id=" + merchantId + "&merchant_user_id=" + merchantUserId + "&service_id=" + serviceId + "&merchant_trans_id=" + merchantTransId + "&return_url" + returnUrl
	return map[string]interface{}{
		"ID":   transaction.ID,
		"Logo": logo,
		"Link": link,
	}
}
