package octo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getUUID(data interface{}) (OctoGetUUIDResponse, error) {
	jsonData, _ := json.Marshal(data)
	client := &http.Client{}
	url := OCTO_URL + "/prepare_payment"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var octoGetUUIDResponse OctoGetUUIDResponse
	err := json.Unmarshal(body, &octoGetUUIDResponse)
	if err != nil {
		return octoGetUUIDResponse, nil
	}

	return octoGetUUIDResponse, nil
}

func getLink(data interface{}, uuid string) (OctoPrepareGetLinkResponse, error) {
	jsonData, _ := json.Marshal(data)
	url := OCTO_URL + "/pay/" + uuid
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OctoPrepareGetLinkResponse{}, fmt.Errorf("Javobni o‘qishda xatolik: %w", err)
	}

	var result OctoPrepareGetLinkResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return OctoPrepareGetLinkResponse{}, fmt.Errorf("JSON parse xatolik: %w", err)
	}

	return result, nil
}
