package model

type Order struct {
	OrderID  int     `json:"order_id"`
	Items    []Item  `json:"items"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
