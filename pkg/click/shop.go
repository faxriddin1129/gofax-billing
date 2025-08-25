package click

import (
	"gofax-billing/internal/constants"
	"gofax-billing/internal/models"
	"gofax-billing/pkg/main_server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	link := serviceUrl + "?amount=" + amount + "&merchant_id=" + MerchantId + "&merchant_user_id=" + MerchantUserId + "&service_id=" + ServiceId + "&transaction_param=" + merchantTransId + "&merchant_trans_id=" + merchantTransId + "&return_url=" + returnUrl

	return map[string]interface{}{
		"ID":     transaction.ID,
		"Link":   link,
		"Method": "GET",
	}, http.StatusOK, "FastPay successful"
}

func NotifyShopApi(form *ResponseShopApi, c *gin.Context) {
	if form.MerchantPrepareId == 0 {
		prepare(form, c)
		return
	}
	complete(form, c)
	return
}

func prepare(form *ResponseShopApi, c *gin.Context) {

	errorCode := 0
	transactionId, _ := strconv.ParseInt(form.MerchantTransId, 10, 64)

	transactionModel := models.TransactionGetById(transactionId)
	if transactionModel.ID == 0 {
		errorCode = 1
	}

	res := map[string]interface{}{
		"click_trans_id":      form.ClickTransId,
		"merchant_trans_id":   form.MerchantTransId,
		"merchant_prepare_id": form.MerchantTransId,
		"merchant_confirm_id": form.MerchantTransId,
		"error":               errorCode,
		"error_note":          form.ErrorNote,
	}

	c.JSON(http.StatusOK, res)
	return
}

func complete(form *ResponseShopApi, c *gin.Context) {
	errorCode := 0
	transactionId, _ := strconv.ParseInt(form.MerchantTransId, 10, 64)

	transactionModel := models.TransactionGetById(transactionId)
	if transactionModel.ID == 0 {
		errorCode = 1
	}

	if errorCode == 0 {
		transactionModel.Status = constants.STATUS_SUCCESS
		transactionModel.PaymentStatus = 1
		_, _ = models.TransactionUpdate(&transactionModel)

		code, _ := main_server.MainServerStatus(transactionModel)
		if code != 0 {
			errorCode = 1
		}

	}

	res := map[string]interface{}{
		"click_trans_id":      form.ClickTransId,
		"merchant_trans_id":   form.MerchantTransId,
		"merchant_confirm_id": form.MerchantTransId,
		"error":               errorCode,
		"error_note":          form.ErrorNote,
	}

	c.JSON(http.StatusOK, res)
	return
}
