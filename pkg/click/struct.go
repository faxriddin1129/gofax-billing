package click

type ResponseShopApi struct {
	ClickTransId      int64   `form:"click_trans_id"`
	ServiceId         int     `form:"service_id"`
	ClickPayDocId     int64   `form:"click_paydoc_id"`
	MerchantTransId   string  `form:"merchant_trans_id"`
	Amount            float64 `form:"amount"`
	Error             int     `form:"error"`
	ErrorNote         string  `form:"error_note"`
	MerchantPrepareId int     `form:"merchant_prepare_id,omitempty"`
}
