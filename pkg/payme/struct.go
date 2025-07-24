package payme

type PaymeRequest struct {
	Method string       `json:"method"`
	Params *PaymeParams `json:"params"`
	ID     int64        `json:"id,omitempty"`
}

type PaymeParams struct {
	Amount  float64  `json:"amount,omitempty"`
	ID      string   `json:"id,omitempty"`
	Time    int64    `json:"time,omitempty"`
	Reason  int      `json:"reason,omitempty"`
	From    int64    `json:"from,omitempty"`
	To      int64    `json:"to,omitempty"`
	Account *Account `json:"account,omitempty"`
}

type Account struct {
	OrderId string `json:"order_id,omitempty"`
}

type ErrorMessage struct {
	Ru string `json:"ru,omitempty"`
	Uz string `json:"uz,omitempty"`
	En string `json:"en,omitempty"`
}

type ErrorResponse struct {
	Error struct {
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
	} `json:"error"`
}

type SuccessResponse struct {
	Result struct {
		Allow bool `json:"allow"`
	} `json:"result"`
}

type CanceledResponse struct {
	Result struct {
		Transaction string `json:"transaction"`
		State       int    `json:"state"`
		CancelTime  int64  `json:"cancel_time"`
		CreateTime  int64  `json:"create_time"`
		PerformTime int64  `json:"perform_time"`
		Reason      int    `json:"reason"`
	} `json:"result"`
}
