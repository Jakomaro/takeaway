package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ValidateGetMethod(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)

			errMsg := fmt.Sprintf("method %s not allowed", r.Method)
			errMap := map[string]string{"error": errMsg}

			err := json.NewEncoder(w).Encode(errMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateBody(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Contest-Type") != "application/json" {
			http.Error(w, "failed to validate Contest-Type", http.StatusBadRequest)
			return
		}

		if r.Body == nil {
			http.Error(w, "failed to validate the body", http.StatusBadRequest)
		}

		next.ServeHTTP(w, r)
	})

}
