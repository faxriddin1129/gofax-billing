package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const Token = "6948992442:AAE2CzY1eP4JvnB5qE9qz6bl1OpaBce2p_w"

func SendTelegramMessage(data map[string]interface{}) (map[string]interface{}, error) {
	url := "https://api.telegram.org/bot" + Token + "/sendMessage"

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
