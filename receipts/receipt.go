package receipts

import "github.com/programcpp/receipt-processor/json_time"

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            float32 `json:"price,string"`
}
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate json_time.Date `json:"purchaseDate"`
	PurchaseTime json_time.Time `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        float32 `json:"total,string"`
}
