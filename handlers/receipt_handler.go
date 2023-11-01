package handlers

import "net/http"

func CreateReceipts(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
}