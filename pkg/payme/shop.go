package payme

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"gofax-billing/internal/constants"
	"gofax-billing/internal/models"
	"gofax-billing/pkg/main_server"
	"gofax-billing/pkg/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var MerchantId string

func GenerateShopApiLink(transaction *models.Transaction) (interface{}, int, string) {

	if transaction.Platform == constants.PlatformHikmat {
		MerchantId = HIKMAT_MERCHANT_ID
	} else if transaction.Platform == constants.PlatformAsia {
		MerchantId = ASIA_MERCHANT_ID
	}

	serviceUrl := SERVICE_URL
	amount := strconv.FormatFloat(transaction.Amount, 'f', -1, 64)
	orderId := strconv.Itoa(int(transaction.ID))
	returnUrl := transaction.ReturnUrl
	language := "ru"

	paramStr := "m=" + MerchantId + ";ac.order_id=" + orderId + ";a=" + amount + ";c=" + returnUrl + ";l=" + language
	encoded := base64.StdEncoding.EncodeToString([]byte(paramStr))
	link := serviceUrl + "/" + encoded

	return map[string]interface{}{
		"ID":     transaction.ID,
		"Link":   link,
		"Method": "GET",
	}, http.StatusOK, "FastPay successful"
}

func NotifyShopApi(form *PaymeRequest, c *gin.Context) {

	if !CheckAuthHeader(c) {
		c.JSON(http.StatusOK, NoAuth())
		return
	}

	switch form.Method {
	case "CheckPerformTransaction":
		CheckPerformTransaction(form, c)
	case "CreateTransaction":
		CreateTransaction(form, c)
	case "PerformTransaction":
		PerformTransaction(form, c)
	case "CancelTransaction":
		CancelTransaction(form, c)
	case "CheckTransaction":
		CheckTransaction(form, c)
	case "GetStatement":
		GetStatement(form, c)
	default:
		c.JSON(http.StatusOK, NotFound())
	}
}

func CheckPerformTransaction(form *PaymeRequest, c *gin.Context) {

	if form.Params.Account.OrderId == "" {
		c.JSON(http.StatusOK, NotParam())
		return
	}

	orderID, _ := strconv.ParseInt(form.Params.Account.OrderId, 10, 64)
	if orderID == 0 {
		c.JSON(http.StatusOK, NotParam())
		return
	}

	transaction := models.TransactionGetById(orderID)
	if transaction.ID == 0 {
		c.JSON(http.StatusOK, NotParam())
		return
	}

	if transaction.PaymentStatus == 1 {
		c.JSON(http.StatusOK, NotFound())
		return
	}

	if transaction.State == -2 {
		c.JSON(http.StatusOK, Canceled(&transaction))
		return
	}

	if form.Params.Amount != transaction.Amount {
		c.JSON(http.StatusOK, NotCorrectAmount())
		return
	}

	transaction.CreateTime = 0
	transaction.UUID = ""
	transaction.PerformTime = 0
	transaction.CancelTime = 0
	transaction.Status = 0
	transaction.Reason = 0
	_, err := models.TransactionUpdate(&transaction)

	fmt.Println(err)

	c.JSON(http.StatusOK, Success())
	return
}

func CreateTransaction(form *PaymeRequest, c *gin.Context) {
	timeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	uuid := form.Params.ID
	orderID, _ := strconv.ParseInt(form.Params.Account.OrderId, 10, 64)
	if orderID == 0 {
		c.JSON(http.StatusOK, NotParam())
		return
	}

	transaction := models.TransactionGetById(orderID)

	if transaction.ID == 0 {
		c.JSON(http.StatusOK, NotParam())
		return
	}

	if transaction.Amount != form.Params.Amount {
		c.JSON(http.StatusOK, NotCorrectAmount())
		return
	}
	if transaction.State == -2 {
		c.JSON(http.StatusOK, Canceled(&transaction))
		return
	}
	if transaction.State == 1 && transaction.UUID != uuid {
		c.JSON(http.StatusOK, Pending())
		return
	}

	if transaction.State == 1 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": map[string]interface{}{
				"create_time": transaction.CreateTime,
				"transaction": strconv.Itoa(int(transaction.ID)),
				"state":       transaction.State,
			},
		})
		return
	}

	if transaction.ID == 0 {
		c.JSON(http.StatusOK, NotFound())
		return
	}

	transaction.CreateTime = timeMillis
	transaction.UUID = uuid
	transaction.State = 1
	_, _ = models.TransactionUpdate(&transaction)
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": map[string]interface{}{
			"create_time": timeMillis,
			"transaction": strconv.Itoa(int(transaction.ID)),
			"state":       transaction.State,
		},
	})
}

func PerformTransaction(form *PaymeRequest, c *gin.Context) {

	transaction := models.TransactionGetByUUID(form.Params.ID)
	if transaction.ID == 0 {
		c.JSON(http.StatusOK, NotFound())
		return
	}

	if transaction.State == -2 {
		c.JSON(http.StatusOK, Canceled(&transaction))
		return
	}

	timeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	if transaction.PerformTime == 0 {
		transaction.PerformTime = timeMillis
		transaction.State = 2
		transaction.PaymentStatus = 1
		transaction.Status = constants.STATUS_SUCCESS
		_, _ = models.TransactionUpdate(&transaction)
	}

	code, _ := main_server.MainServerStatus(transaction)
	if code == 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": map[string]interface{}{
				"perform_time": transaction.PerformTime,
				"transaction":  strconv.Itoa(int(transaction.ID)),
				"state":        transaction.State,
			},
		})
		return
	}
	c.JSON(http.StatusOK, NotParam())
	return
}

func CheckTransaction(form *PaymeRequest, c *gin.Context) {

	trId := form.Params.ID
	transaction := models.TransactionGetByUUID(trId)

	if transaction.ID != 0 {
		if transaction.Status == -2 {
			c.JSON(http.StatusOK, Canceled(&transaction))
			return
		} else {
			var reason interface{}
			reason = transaction.Reason
			if transaction.Reason == 0 {
				reason = nil
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"result": map[string]interface{}{
					"create_time":  transaction.CreateTime,
					"perform_time": transaction.PerformTime,
					"cancel_time":  transaction.CancelTime,
					"transaction":  strconv.Itoa(int(transaction.ID)),
					"state":        transaction.State,
					"reason":       reason,
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, NotFound())
}

func CancelTransaction(form *PaymeRequest, c *gin.Context) {
	timeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	uuid := form.Params.ID
	transaction := models.TransactionGetByUUID(uuid)
	if transaction.ID != 0 {
		if transaction.CancelTime == 0 {
			if transaction.PaymentStatus == 0 {
				transaction.State = -1
			} else {
				transaction.State = -2
			}
			transaction.CancelTime = timeMillis
			transaction.Reason = form.Params.Reason
			transaction.Status = constants.STATUS_CANCEL
			res, err := models.TransactionUpdate(&transaction)
			fmt.Println(res)
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": map[string]interface{}{
				"cancel_time": transaction.CancelTime,
				"transaction": strconv.Itoa(int(transaction.ID)),
				"state":       transaction.State,
			},
		})
		return
	}

	c.JSON(http.StatusOK, NotFound())
}

func CheckAuthHeader(c *gin.Context) bool {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return false
	}

	expectedKeyHimat := "Paycom:" + HIKMAT_PROD_KEY
	expectedKeyAsia := "Paycom:" + ASIA_PROD_KEY

	if strings.HasPrefix(strings.ToLower(authHeader), "basic ") {
		encoded := strings.TrimSpace(authHeader[6:])
		decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return false
		}

		decoded := string(decodedBytes)
		if decoded != expectedKeyHimat || decoded != expectedKeyAsia {
			return false
		}

		return true
	}

	return false
}

func GetStatement(form *PaymeRequest, c *gin.Context) {
	var data []models.Transaction
	utils.DB.Where("create_time >= ? AND create_time <= ?", form.Params.From, form.Params.To).Find(&data)

	var result []map[string]interface{}
	for _, v := range data {

		var state interface{}
		state = v.State
		if state == 0 {
			state = nil
		}

		var reason interface{}
		reason = v.Reason
		if reason == 0 {
			reason = nil
		}

		result = append(result, map[string]interface{}{
			"id":     v.ID,
			"time":   v.CreatedAt,
			"amount": v.Amount,
			"account": map[string]string{
				"order_id": v.OrderId,
			},
			"create_time":  v.CreateTime,
			"perform_time": v.PerformTime,
			"cancel_time":  v.CancelTime,
			"transaction":  v.TransactionId,
			"state":        state,
			"reason":       reason,
			"receivers": []map[string]interface{}{
				{
					"id":     v.ID,
					"amount": v.Amount,
				},
			},
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": map[string]interface{}{
			"transactions": result,
		},
	})
}
