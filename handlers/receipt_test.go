package handlers_test

import (
	"encoding/json"
	"testing"

	"github.com/programcpp/receipt-processor/handlers"
	"github.com/stretchr/testify/assert"
)

func TestReceiptJsonCodec(t *testing.T) {
	receiptStr := `{
		"retailer": "abc",
		"purchaseDate": "2023-11-01",
		"purchaseTime": "23:30",
		"items": [
		  {
			"shortDescription": "item 1 des",
			"price": "10.50"
		  }
		],
		"total": "10.50"
	}`

	receipt := handlers.Receipt{}
	err := json.Unmarshal([]byte(receiptStr), &receipt)
	assert.NoError(t, err)

	assert.Equal(t, "abc", receipt.Retailer)
	assert.Equal(t, "2023-11-01", receipt.PurchaseDate)
	assert.Equal(t, "23:30", receipt.PurchaseTime)
	assert.Equal(t, 1, len(receipt.Items))
	assert.Equal(t, "item 1 des", receipt.Items[0].ShortDescription)
	assert.Equal(t, "10.50", receipt.Items[0].Price)
}
