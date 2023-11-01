package main

import (
	"log"
	"net/http"

	"github.com/programcpp/receipt-processor/handlers"
)

func main(){
	mux := http.NewServeMux()
	
	// init api handlers
	mux.HandleFunc("/receipts/process", handlers.CreateReceipts)

	log.Println("listening on port 1201...")
	log.Fatal(http.ListenAndServe(":1201", mux))
}