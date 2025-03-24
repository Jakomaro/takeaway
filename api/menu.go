package api

import (
	"encoding/json"
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
	err := json.NewEncoder(w).Encode(h.menuService.GetMenu())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
