package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/google/uuid"
	"github.com/programcpp/receipt-processor/mocks"
	"github.com/programcpp/receipt-processor/handlers"
	"github.com/programcpp/receipt-processor/receipts"
	"github.com/programcpp/receipt-processor/test_utils"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestCreateReceiptSuccess(t *testing.T) {
	mockDb := mocks.NewDb(t)
	handler := handlers.NewHandler(mockDb)

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

	mockResId := uuid.New().String()
	mockDb.On("Put", receipts.Receipt{
		Retailer:     "abc",
		PurchaseDate: test_utils.GetDate("2023-11-01"),
		PurchaseTime: test_utils.GetTime("23:30"),
		Items: []receipts.Item{
			{
				ShortDescription: "item 1 des",
				Price:            10.50,
			},
		},
		Total: 10.50,
	}).Return(mockResId)
	handler.Create(w, req)
	resp := w.Result()

	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, json.Valid(respBody))
	resId := gjson.Get(string(respBody), "id").String()
	_, err = uuid.Parse(resId)
	assert.NoError(t, err)
	assert.Equal(t, mockResId, resId)
}

func TestCreateReceiptFailureOnBadPayload(t *testing.T) {
	mockDb := mocks.NewDb(t)
	handler := handlers.NewHandler(mockDb)

	reqStr := `{
	}`
	req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBufferString(reqStr))
	w := httptest.NewRecorder()

	handler.Create(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
