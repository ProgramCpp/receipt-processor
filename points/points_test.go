package points_test

import (
	"testing"

	"github.com/programcpp/receipt-processor/points"
	"github.com/programcpp/receipt-processor/receipts"
	"github.com/stretchr/testify/assert"
)

func TestPoints(t *testing.T){
	points := points.Points(receipts.Receipt{
		Retailer: "abc",
	})
	assert.Equal(t, 3, points)
}