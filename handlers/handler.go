package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/programcpp/receipt-processor/db"
	"github.com/programcpp/receipt-processor/points"
	"github.com/programcpp/receipt-processor/receipts"
)

type Handler struct {
	// handler directly communicates with persistence layer
	// since the business layer is lean
	database db.Db
}

func NewHandler(d db.Db) Handler {
	return Handler{d}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: validate request before unmarshaling
	// amount values must have exactly 2 decimal places only

	receipt := receipts.Receipt{}
	err = json.Unmarshal(body, &receipt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !receipt.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := h.database.Put(receipt)

	// response is simple - for now - which can be handcoded into json
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id))) // status code is set sutomatically
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if(len(params) != 1){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, ok := params["id"]; 
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	iReceipt, found := h.database.Get(id)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	receipt, ok := iReceipt.(receipts.Receipt)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	points := points.NewEngine().Points(receipt)

	// response is simple - for now - which can be handcoded into json
	w.Write([]byte(fmt.Sprintf("{\"points\":\"%d\"}", points))) // status code is set sutomatically	
}