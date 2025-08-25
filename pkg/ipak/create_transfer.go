package ipak

import (
	"gofax-billing/internal/constants"
	"gofax-billing/internal/models"
	"net/http"
	"strconv"
)

var TOKEN string

func CreateTransfer(transaction *models.Transaction) (interface{}, int, string) {

	if transaction.Platform == constants.PlatformAsia {
		TOKEN = ASIA_TOKEN
	}

	if TOKEN == "" {
		return map[string]interface{}{
			"ID":     0,
			"Link":   nil,
			"Method": nil,
		}, http.StatusBadRequest, "Token not found"
	}

	generateData := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "transfer.create",
		"id":      strconv.Itoa(int(transaction.ID)),
		"params": map[string]interface{}{
			"order_id": transaction.OrderId,
			"amount":   transaction.Amount,
			"card": map[string]interface{}{
				"pan":    transaction.CardNumber,
				"expiry": transaction.CardExpire,
			},
			"details": map[string]interface{}{
				"description": "Travel Payment",
				"ofdInfo": map[string]interface{}{
					"ReceiptType": 0,
					"Items": []interface{}{
						map[string]interface{}{
							"Name":        "Travel Product",
							"SPIC":        "10103002001001002",
							"PackageCode": "10103002001001002",
							"price":       transaction.Amount,
							"count":       1,
							"VATPercent":  0,
							"Discount":    0,
						},
					},
				},
			},
		},
	}

	remote := CreateTransferRequest(generateData, TOKEN)

	if remote.Error != nil && remote.Error.Code != 0 {
		return map[string]interface{}{
			"ID":     0,
			"Link":   nil,
			"Method": nil,
		}, http.StatusBadRequest, remote.Error.Message
	}

	transaction.UUID = remote.Result.TransferId
	_, _ = models.TransactionUpdate(transaction)

	return map[string]interface{}{
		"ID":     transaction.ID,
		"Link":   nil,
		"Method": nil,
	}, http.StatusOK, "Conformation code send to your phone"
}
