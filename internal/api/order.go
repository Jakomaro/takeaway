package api

import "net/http"

func PostOrder(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
}
