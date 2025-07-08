package octo

type OctoGetUUIDResponse struct {
	Error                   int              `json:"error"`
	ErrMessage              string           `json:"errMessage,omitempty"`
	OctoPaymentUUID         string           `json:"octo_payment_UUID"`
	ApiMessageForDevelopers string           `json:"apiMessageForDevelopers,omitempty"`
	Data                    *OctoGetUUIDData `json:"data,omitempty"`
}

type OctoGetUUIDData struct {
	ShopTransactionID string `json:"shop_transaction_id"`
	OctoPaymentUUID   string `json:"octo_payment_UUID"`
	OctoPayURL        string `json:"octo_pay_url"`
}

type OctoPrepareGetLinkResponse struct {
	Error                   int                     `json:"error"`
	ErrMessage              string                  `json:"errMessage"`
	ApiMessageForDevelopers string                  `json:"apiMessageForDevelopers,omitempty"`
	Data                    *OctoPrepareGetLinkData `json:"data,omitempty"`
}

type OctoPrepareGetLinkData struct {
	UUID        string `json:"uuid"`
	RedirectURL string `json:"redirectUrl"`
}

type OctoNotifyResponse struct {
	ShopTransactionID string `json:"shop_transaction_id"`
	OctoPaymentUUID   string `json:"octo_payment_UUID"`
	Status            string `json:"status"`
	Signature         string `json:"signature"`
	HashKey           string `json:"hash_key"`
}
