package receipts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/programcpp/receipt-processor/db"
)

type Handler struct {
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
	receipt := Receipt{}
	err = json.Unmarshal(body, &receipt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := h.database.Insert(receipt)

	// response is simple - for now - which can be handcoded into json
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id))) // status code is set sutomatically
}
