package model

type Item struct {
	ItemID int     `json:"item_id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
}
