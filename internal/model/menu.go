package model

type Item struct {
	ItemID int     `json:"item_id" db:"item_id"`
	Name   string  `json:"name" db:"name"`
	Price  float64 `json:"price" db:"price"`
}
