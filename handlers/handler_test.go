package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/programcpp/receipt-processor/handlers"
	"github.com/programcpp/receipt-processor/mocks"
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

func TestGetPointsFailureOnBadRequest(t *testing.T) {
	mockDb := mocks.NewDb(t)
	handler := handlers.NewHandler(mockDb)

	// id not present
	req := httptest.NewRequest("GET", "/receipts//points", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// not a uuid
	req = httptest.NewRequest("GET", "/receipts/123/points", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	w = httptest.NewRecorder()

	handler.Get(w, req)
	resp = w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetPointsFailureReceiptNotFound(t *testing.T) {
	mockDb := mocks.NewDb(t)
	handler := handlers.NewHandler(mockDb)

	id := uuid.New().String()
	req, err := http.NewRequest("GET", "/receipts/"+id+"/points", nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": id})
	w := httptest.NewRecorder()

	mockDb.On("Get", id).Return("", false)
	handler.Get(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetPointsFailureOnCorruptedDbEntry(t *testing.T) {
	mockDb := mocks.NewDb(t)
	handler := handlers.NewHandler(mockDb)

	id := uuid.New().String()
	req, err := http.NewRequest("GET", "/receipts/"+id+"/points", nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": id})
	w := httptest.NewRecorder()

	mockDb.On("Get", id).Return("not a receipt", true)
	handler.Get(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestGetPointsSuccess(t *testing.T) {
	mockDb := mocks.NewDb(t)
	handler := handlers.NewHandler(mockDb)

	id := uuid.New().String()
	req, err := http.NewRequest("GET", "/receipts/"+id+"/points", nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": id})
	w := httptest.NewRecorder()

	mockDb.On("Get", id).Return(receipts.Receipt{
		Retailer:     "walmart",
		PurchaseDate: test_utils.GetDate("2023-11-01"),
		PurchaseTime: test_utils.GetTime("02:30"),
		Items: []receipts.Item{
			{
				ShortDescription: "coke",
				Price:            2.49,
			},
			{
				ShortDescription: "cake",
				Price:            14.99,
			},
			{
				ShortDescription: "swiss roll",
				Price:            5.00,
			},
		},
		Total: 10.0,
	}, true)
	handler.Get(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t, int64(93), gjson.GetBytes(resBody, "points").Int())
}
