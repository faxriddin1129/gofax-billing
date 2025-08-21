package main_server

import (
	"bytes"
	"encoding/json"
	"gofax-billing/internal/models"
	"io"
	"net/http"
)

type Notify struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func MainServerStatus(transaction models.Transaction) (int, string) {
	url := MAIN_SERVER_URL + "?token=" + MAIN_SERVER_TOKEN

	jsonData, err := json.Marshal(transaction)
	if err != nil {
		return 1, "Error"
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 1, "Error"
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 1, "Error"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 1, "Error"
	}

	var notify Notify
	if err := json.Unmarshal(body, &notify); err != nil {
		return 1, "Error"
	}

	return notify.Code, notify.Message
}
