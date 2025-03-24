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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600") //1 day

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(h.menuService.GetMenu())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
