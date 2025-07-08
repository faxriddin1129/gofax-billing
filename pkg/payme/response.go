package payme

import (
	"microservice/internal/models"
	"strconv"
)

func Success() SuccessResponse {
	return SuccessResponse{
		Result: struct {
			Allow bool `json:"allow"`
		}{
			Allow: true,
		},
	}
}

func newErrorResponse(code int, message interface{}) ErrorResponse {
	return ErrorResponse{
		Error: struct {
			Code    int         `json:"code"`
			Message interface{} `json:"message"`
		}{
			Code:    code,
			Message: message,
		},
	}
}

func NotFound() ErrorResponse {
	return newErrorResponse(-31003, ErrorMessage{
		Ru: "Транзакция не найдена",
		Uz: "Tranzaksiya topilmadi",
		En: "Transaction not found",
	})
}

func NotParam() ErrorResponse {
	return newErrorResponse(-31050, ErrorMessage{
		Ru: "Ошибки неверного ввода данных покупателем",
		Uz: "Xaridor tomonidan noto`g`ri ma`lumotlarni kiritish xatolari",
		En: "Errors of incorrect data entry by the buyer",
	})
}

func Pending() ErrorResponse {
	return newErrorResponse(-31050, ErrorMessage{
		Ru: "В ожидании оплаты",
		Uz: "To`lov kutilmoqda",
		En: "Waiting for payment",
	})
}

func NotCorrectAmount() ErrorResponse {
	return newErrorResponse(-31001, ErrorMessage{
		Ru: "Неверная сумма",
		Uz: "Yaroqsiz miqdor",
		En: "Incorrect amount",
	})
}

func NoAuth() ErrorResponse {
	return newErrorResponse(-32504, "Недостаточно привилегий для выполнения метода")
}

func Canceled(tx *models.Transaction) CanceledResponse {
	return CanceledResponse{
		Result: struct {
			Transaction string `json:"transaction"`
			State       int    `json:"state"`
			CancelTime  int64  `json:"cancel_time"`
			CreateTime  int64  `json:"create_time"`
			PerformTime int64  `json:"perform_time"`
			Reason      int    `json:"reason"`
		}{
			Transaction: strconv.Itoa(int(tx.ID)),
			State:       int(tx.State),
			CancelTime:  tx.CancelTime,
			CreateTime:  tx.CreateTime,
			PerformTime: tx.PerformTime,
			Reason:      tx.Reason,
		},
	}
}
