package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReceiptsSuccess(t *testing.T){
	router := newServer()
	server := httptest.NewServer(router)
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

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, json.Valid(resBody))
	assert.Contains(t, string(resBody), `"id"`)
}

func TestCreateReceiptsFailure_InvalidMethod(t *testing.T){
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, nil)
	assert.NoError(t, err)
	res, err := server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)


	req, err = http.NewRequest("PUT", server.URL, nil)
	assert.NoError(t, err)
	res, err = server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)


	req, err = http.NewRequest("PATCH", server.URL, nil)
	assert.NoError(t, err)
	res, err = server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}