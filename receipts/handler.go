package receipts

import (
	"encoding/json"
	"io"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
}
