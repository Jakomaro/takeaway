package api

// import (
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/jakomaro/takeaway/internal/model"
// )

// func PostOrder(w http.ResponseWriter, r *http.Request) {

// 	//TODO validate request needs to go in a separate middleware function
// 	body := validateBody(r, w)
// 	if body == nil {
// 		return
// 	}
// 	var order model.OrderRequest
// 	err := json.Unmarshal(body, &order)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func validateBody(r *http.Request, w http.ResponseWriter) []byte {
// 	if r.Header.Get("Content-Type") != "application/json" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return nil
// 	}
// 	if r.Body == nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return nil
// 	}

// 	r.Body = http.MaxBytesReader(w, r.Body, 1024*100)
// 	body, err := io.ReadAll(r.Body)
// 	defer r.Body.Close()
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return nil
// 	}
// 	if len(body) == 0 {
// 		log.Println("empty body")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return nil
// 	}
// 	return body
// }
