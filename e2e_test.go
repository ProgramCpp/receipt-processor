package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReceiptsSuccess(t *testing.T){
	mux := newServer()
	server := httptest.NewServer(mux)
	defer server.Close()
	
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
	res, err := http.Post(server.URL + "/receipts/process", "application/json", bytes.NewBufferString(reqStr))
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}