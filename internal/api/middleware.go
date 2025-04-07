package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
			log.Println("Error: ValidateBody - missing Contest-Type")
			http.Error(w, "failed to validate Contest-Type", http.StatusBadRequest)
			return
		}

		if r.Body == nil {
			log.Println("Error: ValidateBody - missing body")
			http.Error(w, "failed to validate the body", http.StatusBadRequest)
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1024*100)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error: ValidateBody - read body: %v", err)
			http.Error(w, "failed to read body", http.StatusBadRequest)
		}
		defer r.Body.Close()
		if len(body) == 0 {
			log.Println("Error: ValidateBody - empty body")
			http.Error(w, "failed to read body", http.StatusBadRequest)
		}
		// copy back the body into the request
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})

}
