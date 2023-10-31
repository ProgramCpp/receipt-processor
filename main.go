package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	log.Fatal(http.ListenAndServe(":1201", mux))
}