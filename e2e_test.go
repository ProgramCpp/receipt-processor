package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReceiptsSuccess(t *testing.T){
	mux := newServer()
	server := httptest.NewServer(mux)
	defer server.Close()
	
	res, err := http.Post(server.URL + "/receipts/process", "application/json", nil)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}