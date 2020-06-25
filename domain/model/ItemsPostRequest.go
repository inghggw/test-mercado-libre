package model

// ItemsPostRequest struct
type ItemsPostRequest struct {
	ItemIDs []string `json:"item_ids"`
	Amount  float64 `json:"amount"`
}