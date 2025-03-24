package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jakomaro/takeaway/services"
)

type MenuHandler struct {
	menuService services.MenuServicer
}

func NewMenuHandler(menuService services.MenuServicer) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

func (h *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errMsg := fmt.Sprintf("method %s not allowed", r.Method)
		errMap := map[string]string{"error": errMsg}

		err := json.NewEncoder(w).Encode(errMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := json.NewEncoder(w).Encode(h.menuService.GetMenu())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
