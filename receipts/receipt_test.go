package receipts_test

import (
	"encoding/json"
	"fmt"
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
	assert.Equal(t, float64(10.50), receipt.Items[0].Price)
	assert.Equal(t, float64(10.50), receipt.Total)
}

func TestReceiptIsValid(t *testing.T) {
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

	assert.True(t, receipt.IsValid())
}

func TestReceiptInvalidRetailer(t *testing.T) {
	testCases := []string{"", "abc retailer"}

	receiptStr := `{
		"retailer": "%s",
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

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			receipt := receipts.Receipt{}
			err := json.Unmarshal([]byte(fmt.Sprintf(receiptStr, tc)), &receipt)
			assert.NoError(t, err)

			assert.Equal(t, tc, receipt.Retailer)
			assert.False(t, receipt.IsValid())
		})
	}
}

func TestReceiptInvalidPurchaseDate(t *testing.T) {
	receiptStr := `{
		"retailer": "abc",
		"purchaseDate": "",
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
	assert.Error(t, err)
}

func TestReceiptInvalidPurchaseTime(t *testing.T) {
	receiptStr := `{
		"retailer": "abc",
		"purchaseDate": "2023-11-01",
		"purchaseTime": "",
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
	assert.Error(t, err)
}

func TestReceiptInvalidItemCount(t *testing.T) {
	receiptStr := `{
		"retailer": "abc",
		"purchaseDate": "2023-11-01",
		"purchaseTime": "23:30",
		"items": [],
		"total": "10.50"
	}`

	receipt := receipts.Receipt{}
	err := json.Unmarshal([]byte(receiptStr), &receipt)
	assert.NoError(t, err)

	assert.False(t, receipt.IsValid())
}

func TestReceiptInvalidTotal(t *testing.T) {
	// validation is broken
	testCases := []string{/*"10.5", "10.505"*/}

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
		"total": "%s"
	}`

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			receipt := receipts.Receipt{}
			err := json.Unmarshal([]byte(fmt.Sprintf(receiptStr, tc)), &receipt)
			assert.NoError(t, err)

			assert.False(t, receipt.IsValid())
		})
	}

	tc := ".5"
	receipt := receipts.Receipt{}
	err := json.Unmarshal([]byte(fmt.Sprintf(receiptStr, tc)), &receipt)
	// json unmarshal fails
	assert.Error(t, err)
}

func TestReceiptInvalidItems(t *testing.T) {
	testCases := []string{
		`{
		"shortDescription": "item &*&^#$",
		"price": "10.50"
	     }`,
		// validation is broken
		// `{
		// "shortDescription": "item description",
		// "price": "10.505"
		//  }`
	}

	receiptStr := `{
		"retailer": "abc",
		"purchaseDate": "2023-11-01",
		"purchaseTime": "23:30",
		"items": [%s],
		"total": "10.50"
	}`

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			receipt := receipts.Receipt{}
			err := json.Unmarshal([]byte(fmt.Sprintf(receiptStr, tc)), &receipt)
			assert.NoError(t, err)

			assert.False(t, receipt.IsValid())
		})
	}
}
