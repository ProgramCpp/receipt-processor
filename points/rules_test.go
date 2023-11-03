package points_test

import (
	"testing"
	"time"

	"github.com/programcpp/receipt-processor/json_time"
	"github.com/programcpp/receipt-processor/points"
	"github.com/programcpp/receipt-processor/receipts"
	"github.com/programcpp/receipt-processor/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestRetailerRule(t *testing.T) {
	assert.Equal(t, 3, points.RetailerRule(receipts.Receipt{Retailer: "abc"}))
	assert.Equal(t, 4, points.RetailerRule(receipts.Receipt{Retailer: "abcd"}))
	assert.Equal(t, 4, points.RetailerRule(receipts.Receipt{Retailer: "abc1"}))
	assert.Equal(t, 3, points.RetailerRule(receipts.Receipt{Retailer: "abc#$%"}))
	assert.Equal(t, 3, points.RetailerRule(receipts.Receipt{Retailer: "ab c"}))
}

func TestRoundTotalRule(t *testing.T) {
	assert.Equal(t, 50, points.RoundTotalRule(receipts.Receipt{Total: 50}))
	assert.Equal(t, 0, points.RetailerRule(receipts.Receipt{Total: 50.30}))
}

func TestMultipleTotalRule(t *testing.T) {
	assert.Equal(t, 25, points.MultipleTotalRule(receipts.Receipt{Total: 50}))
	assert.Equal(t, 25, points.MultipleTotalRule(receipts.Receipt{Total: 50.25}))
	assert.Equal(t, 25, points.MultipleTotalRule(receipts.Receipt{Total: 50.50}))
	assert.Equal(t, 25, points.MultipleTotalRule(receipts.Receipt{Total: 50.75}))
	assert.Equal(t, 0, points.MultipleTotalRule(receipts.Receipt{Total: 50.30}))
}
func TestItemCountRule(t *testing.T) {
	assert.Equal(t, 0, points.ItemCountRule(receipts.Receipt{Items: []receipts.Item{{}}}))
	assert.Equal(t, 5, points.ItemCountRule(receipts.Receipt{Items: []receipts.Item{{}, {}}}))
	assert.Equal(t, 5, points.ItemCountRule(receipts.Receipt{Items: []receipts.Item{{}, {}, {}}}))
	assert.Equal(t, 10, points.ItemCountRule(receipts.Receipt{Items: []receipts.Item{{}, {}, {}, {}}}))
}

func TestItemDescriptionRule(t *testing.T) {
	type testCase struct {
		name      string
		inReceipt receipts.Receipt
		outPoints int
	}

	testCases := []testCase{
		{
			name: "a",
			inReceipt: receipts.Receipt{
				Items: []receipts.Item{
					{
						ShortDescription: "a",
						Price:            2.50,
					},
				},
			},
			outPoints: 0,
		},
		// roundUp(2.5 * 0.2)
		{
			name: "abc",
			inReceipt: receipts.Receipt{
				Items: []receipts.Item{
					{
						ShortDescription: "abc",
						Price:            2.50,
					},
				},
			},
			outPoints: 1,
		},
		{
			name: "abcd",
			inReceipt: receipts.Receipt{
				Items: []receipts.Item{
					{
						ShortDescription: "abcd",
						Price:            2.50,
					},
				},
			},
			outPoints: 0,
		},
		{
			name: "abcdef",
			inReceipt: receipts.Receipt{
				Items: []receipts.Item{
					{
						ShortDescription: "abcdef",
						Price:            2.50,
					},
				},
			},
			outPoints: 1,
		},
		{
			name: "multiple items",
			inReceipt: receipts.Receipt{
				Items: []receipts.Item{
					{
						ShortDescription: "abcdef",
						Price:            2.50,
					},
					{
						ShortDescription: "abc",
						Price:            2.50,
					},
				},
			},
			outPoints: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.outPoints, points.ItemDescriptionRule(tc.inReceipt))
		})
	}
}

func TestPurchaseDateRule(t *testing.T) {
	// timezone doesnt matter
	assert.Equal(t, 6, points.PurchaseDateRule(receipts.Receipt{PurchaseDate: json_time.Date{Time: time.Date(0, 0, 1, 0, 0, 0, 0, time.UTC)}}))
	assert.Equal(t, 0, points.PurchaseDateRule(receipts.Receipt{PurchaseDate: json_time.Date{Time: time.Date(0, 0, 2, 0, 0, 0, 0, time.UTC)}}))
	assert.Equal(t, 6, points.PurchaseDateRule(receipts.Receipt{PurchaseDate: json_time.Date{Time: time.Date(0, 0, 3, 0, 0, 0, 0, time.UTC)}}))
	assert.Equal(t, 0, points.PurchaseDateRule(receipts.Receipt{PurchaseDate: json_time.Date{Time: time.Date(0, 0, 4, 0, 0, 0, 0, time.UTC)}}))
}

func TestPurchaseTimeRule(t *testing.T) {
	// timezone doesnt matter
	assert.Equal(t, 0, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("00:30")}))
	assert.Equal(t, 0, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("10:30")}))
	assert.Equal(t, 0, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("13:30")}))
	assert.Equal(t, 0, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("14:00")}))
	assert.Equal(t, 10, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("14:01")}))
	assert.Equal(t, 10, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("14:30")}))
	assert.Equal(t, 10, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("15:00")}))
	assert.Equal(t, 10, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("15:30")}))
	assert.Equal(t, 10, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("15:59")}))
	assert.Equal(t, 0, points.PurchaseTimeRule(receipts.Receipt{PurchaseTime: test_utils.GetTime("16:00")}))
}
