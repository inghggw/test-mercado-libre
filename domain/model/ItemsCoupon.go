package model

// ItemsCoupon model
type ItemsCoupon struct {
	ItemIDs []string `json:"item_ids"`
	Total   float64 `json:"total"`
	Message string `json:"message"`
}
