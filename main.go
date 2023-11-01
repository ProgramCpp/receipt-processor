package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/programcpp/receipt-processor/db"
	"github.com/programcpp/receipt-processor/receipts"
)

func main() {
	mux := newServer()

	log.Println("listening on port 1201...")
	log.Fatal(http.ListenAndServe(":1201", mux))
}

func newServer() http.Handler {
	router := mux.NewRouter()

	d := db.NewMemDb()
	handler := receipts.NewHandler(d)
	// init api handlers
	router.HandleFunc("/receipts/process", handler.Create).Methods("POST")

	return router
}
