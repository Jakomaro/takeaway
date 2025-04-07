package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jakomaro/takeaway/internal/model"
)

func PostOrder(w http.ResponseWriter, r *http.Request) {

	var order model.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
