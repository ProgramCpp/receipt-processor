package main

import (
	"log"
	"net/http"

	"github.com/programcpp/receipt-processor/db"
	"github.com/programcpp/receipt-processor/receipts"
)

func main(){
	mux := newServer()

	log.Println("listening on port 1201...")
	log.Fatal(http.ListenAndServe(":1201", mux))
}

func newServer() http.Handler {
	mux := http.NewServeMux()

	d :=  db.NewMemDb()
	handler := receipts.NewHandler(d)
	// init api handlers
	mux.HandleFunc("/receipts/process", handler.Create)// TODO: add http method

	return mux
}