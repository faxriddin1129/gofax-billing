package ipak

type CreateTransferResponse struct {
	Id     int                   `json:"id,omitempty"`
	Result *CreateTransferResult `json:"result,omitempty"`
	Error  *CreateTransferError  `json:"error,omitempty"`
}

type CreateTransferResult struct {
	TransferId string `json:"transfer_id,omitempty"`
	Code       int    `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

type CreateTransferError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
