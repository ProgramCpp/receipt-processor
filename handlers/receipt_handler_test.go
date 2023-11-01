package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/programcpp/receipt-processor/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateReceiptsSuccess(t *testing.T){
	req := httptest.NewRequest("POST", "http://localhost.com/receipts", nil)
	w := httptest.NewRecorder()
	handlers.CreateReceipts(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}