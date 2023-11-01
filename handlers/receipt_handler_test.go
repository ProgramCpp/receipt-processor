package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/programcpp/receipt-processor/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateReceiptSuccess(t *testing.T) {
	reqStr := `{
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

	req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBufferString(reqStr))
	w := httptest.NewRecorder()
	handlers.CreateReceipts(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
