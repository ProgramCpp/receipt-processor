package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

// TODO: clear test db state after each test
func TestCreateReceiptSuccess(t *testing.T) {
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
	res, err := http.Post(server.URL+"/receipts/process", "application/json", bytes.NewBufferString(reqStr))
	assert.NoError(t, err)

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, json.Valid(resBody))
	id := string(resBody)
	_ = id
	_, err = uuid.Parse(gjson.Get(string(resBody), "id").String())
	assert.NoError(t, err)
}

func TestCreateReceiptFailure_InvalidMethod(t *testing.T) {
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/receipts/process", nil)
	assert.NoError(t, err)
	res, err := server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)

	req, err = http.NewRequest("PUT", server.URL+"/receipts/process", nil)
	assert.NoError(t, err)
	res, err = server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)

	req, err = http.NewRequest("PATCH", server.URL+"/receipts/process", nil)
	assert.NoError(t, err)
	res, err = server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
}

func TestCreateReceiptFailureInvalidReceipt(t *testing.T) {
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	reqStr := `{
		"retailer": "abc (*&^())",
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
	res, err := http.Post(server.URL+"/receipts/process", "application/json", bytes.NewBufferString(reqStr))
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetReceiptFailure_InvalidMethod(t *testing.T) {
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	req, err := http.NewRequest("POST", server.URL+"/receipts/123/points", nil)
	assert.NoError(t, err)
	res, err := server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)

	req, err = http.NewRequest("PUT", server.URL+"/receipts/123/points", nil)
	assert.NoError(t, err)
	res, err = server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)

	req, err = http.NewRequest("PATCH", server.URL+"/receipts/123/points", nil)
	assert.NoError(t, err)
	res, err = server.Client().Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
}

func TestGetPointsSuccess(t *testing.T) {
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	reqStr := `{
		"retailer": "Target",
  		"purchaseDate": "2022-01-01",
  		"purchaseTime": "13:01",
  		"items": [
    		{
      			"shortDescription": "Mountain Dew 12PK",
      			"price": "6.49"
    		},{
      			"shortDescription": "Emils Cheese Pizza",
      			"price": "12.25"
    		},{
      			"shortDescription": "Knorr Creamy Chicken",
      			"price": "1.26"
    		},{
      			"shortDescription": "Doritos Nacho Cheese",
      			"price": "3.35"
    		},{
      			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
     			"price": "12.00"
    		}
  		],
  		"total": "35.35"
	}`
	res, err := http.Post(server.URL+"/receipts/process", "application/json", bytes.NewBufferString(reqStr))
	assert.NoError(t, err)

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	id := string(gjson.GetBytes(resBody, "id").String())

	res, err = http.Get(server.URL + "/receipts/" + id + "/points")
	assert.NoError(t, err)

	resBody, err = io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	assert.Equal(t, int64(28), gjson.GetBytes(resBody, "points").Int())
}

func TestGetPointsSuccess_2(t *testing.T) {
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	// NOTE: retailer name cannot contain spaces as per API spec
	reqStr := `{
		"retailer": "M&M_Corner_Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [
			{
				"shortDescription": "Gatorade",
				"price": "2.25"
			},{
				"shortDescription": "Gatorade",
				"price": "2.25"
			},{
				"shortDescription": "Gatorade",
				"price": "2.25"
			},{
				"shortDescription": "Gatorade",
				"price": "2.25"
			}
		],
		"total": "9.00"
	}`
	res, err := http.Post(server.URL+"/receipts/process", "application/json", bytes.NewBufferString(reqStr))
	assert.NoError(t, err)

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	id := string(gjson.GetBytes(resBody, "id").String())

	res, err = http.Get(server.URL + "/receipts/" + id + "/points")
	assert.NoError(t, err)

	resBody, err = io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	assert.Equal(t, int64(109), gjson.GetBytes(resBody, "points").Int())
}

func TestGetPointsFailureNoReciptFound(t *testing.T) {
	router := newServer()
	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Get(server.URL + "/receipts/" + uuid.New().String() + "/points")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

