package click

import (
	"gofax-billing/internal/constants"
	"gofax-billing/internal/models"
	"net/http"
	"strconv"
)

var MerchantId, MerchantUserId, ServiceId string

func GenerateShopApiLink(transaction *models.Transaction) (interface{}, int, string) {
	amount := strconv.FormatFloat(transaction.Amount, 'f', 2, 64)

	if transaction.Platform == constants.PlatformAsia {
		MerchantId = ASIA_MERCHANT_ID
		MerchantUserId = ASIA_MERCHANT_USER_ID
		ServiceId = ASIA_SERVICE_ID
	} else if transaction.Platform == constants.PlatformHikmat {
		MerchantId = HIKMAT_MERCHANT_ID
		MerchantUserId = HIKMAT_MERCHANT_USER_ID
		ServiceId = HIKMAT_SERVICE_ID
	}

	merchantTransId := strconv.Itoa(int(transaction.ID))
	returnUrl := transaction.ReturnUrl
	serviceUrl := SERVICE_URL

	link := serviceUrl + "?amount=" + amount + "&merchant_id=" + MerchantId + "&merchant_user_id=" + MerchantUserId + "&service_id=" + ServiceId + "&transaction_param=" + merchantTransId + "&return_url=" + returnUrl

	return map[string]interface{}{
		"ID":     transaction.ID,
		"Link":   link,
		"Method": "GET",
	}, http.StatusOK, "FastPay successful"
}
