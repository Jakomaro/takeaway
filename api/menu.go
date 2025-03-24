package api

import (
	"encoding/json"
	"log"
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

	menu, err := h.menuService.GetMenu()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "{\"error\": \"internal server error\"}", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(menu)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "{\"error\": \"internal server error\"}", http.StatusInternalServerError)
		return
	}
}
