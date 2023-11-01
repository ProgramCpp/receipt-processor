package receipts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/programcpp/receipt-processor/db"
)

type Handler struct {
	database *db.MemDb
}

func NewHandler(d *db.MemDb) Handler {
	return Handler{d}
}

type CreateResponse struct {
	Id string `json:"id"`
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

	w.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id)))
}
