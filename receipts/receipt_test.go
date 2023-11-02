package receipts_test

import (
	"encoding/json"
	"testing"

	"github.com/programcpp/receipt-processor/receipts"
	"github.com/programcpp/receipt-processor/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestReceiptJsonCodecSuccess(t *testing.T) {
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

	receipt := receipts.Receipt{}
	err := json.Unmarshal([]byte(receiptStr), &receipt)
	assert.NoError(t, err)

	assert.Equal(t, "abc", receipt.Retailer)
	assert.Equal(t, test_utils.GetDate("2023-11-01"), receipt.PurchaseDate)
	assert.Equal(t, test_utils.GetTime("23:30"), receipt.PurchaseTime)
	assert.Equal(t, 1, len(receipt.Items))
	assert.Equal(t, "item 1 des", receipt.Items[0].ShortDescription)
	assert.Equal(t, float32(10.50), receipt.Items[0].Price)
	assert.Equal(t, float32(10.50), receipt.Total)
}
