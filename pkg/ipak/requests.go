package ipak

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func CreateTransferRequest(data interface{}, TOKEN string) CreateTransferResponse {
	jsonData, _ := json.Marshal(data)
	client := &http.Client{}
	url := IPAK_URL + "/api/transfer"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+TOKEN)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var res CreateTransferResponse
	err := json.Unmarshal(body, &res)
	if err != nil {
		return res
	}

	return res
}
