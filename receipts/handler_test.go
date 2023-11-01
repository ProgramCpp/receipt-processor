package receipts_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/google/uuid"
	"github.com/programcpp/receipt-processor/db"
	"github.com/programcpp/receipt-processor/receipts"
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

	d := db.NewMemDb()
	handler := receipts.NewHandler(d)

	req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBufferString(reqStr))
	w := httptest.NewRecorder()
	handler.Create(w, req)
	resp := w.Result()

	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t, err)

	createRes := receipts.CreateResponse{}
	err = json.Unmarshal(respBody, &createRes)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	_, err = uuid.Parse(createRes.Id)
	assert.NoError(t, err)
}
