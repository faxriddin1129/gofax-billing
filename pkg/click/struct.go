package click

type ResponseShopApi struct {
	ClickTransId      int64   `json:"click_trans_id"`
	ServiceId         int     `json:"service_id"`
	ClickPayDocId     int64   `json:"click_paydoc_id"`
	MerchantTransId   string  `json:"merchant_trans_id"`
	Amount            float64 `json:"amount"`
	Error             int     `json:"error"`
	ErrorNote         string  `json:"error_note"`
	MerchantPrepareId int     `json:"merchant_prepare_id,omitempty"`
}
