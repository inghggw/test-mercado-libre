package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ResponseItem result api
type ResponseItem struct {
	ID     string  `json:"id"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

// GetItemMercadoLibre by itemID
func GetItemMercadoLibre(itemID string) (*ResponseItem, error) {
	var response *ResponseItem
	resp, err := http.Get(os.Getenv("api_mercado_libre") + itemID)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		response := &ResponseItem{}
		response.ID = ""
		response.Price = 0
		response.Status = "404"

		return response, nil
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Printf("Error unmashall: %v\n", response)
		return nil, err
	}

	return response, nil
}
