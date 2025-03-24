package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jakomaro/takeaway/model"
	"github.com/stretchr/testify/assert"
)

type mockMenuService struct {
	menu []model.Item
}

func (m *mockMenuService) GetMenu() []model.Item {
	return m.menu
}

func TestGetMenu(t *testing.T) {

	tests := []struct {
		name       string
		method     string
		tmpMenu    []model.Item
		wantStCode int
		wantBody   string
	}{
		{
			name:       "success empty",
			method:     "GET",
			tmpMenu:    []model.Item{},
			wantStCode: 200,
			wantBody:   "[]",
		},
		{
			name:   "success single item",
			method: "GET",
			tmpMenu: []model.Item{
				{ItemID: 1, Name: "margherita", Price: 4.5},
			},
			wantStCode: 200,
			wantBody:   `[{"item_id":1,"name":"margherita","price":4.5}]`,
		},
		{
			name:   "success multi items",
			method: "GET",
			tmpMenu: []model.Item{
				{ItemID: 1, Name: "focaccia", Price: 5},
				{ItemID: 2, Name: "biancaneve", Price: 6},
				{ItemID: 3, Name: "margherita", Price: 6.5},
			},
			wantStCode: 200,
			wantBody:   `[{"item_id":1,"name":"focaccia","price":5},{"item_id":2,"name":"biancaneve","price":6},{"item_id":3,"name":"margherita","price":6.5}]`,
		},
		{
			name:   "error wrong method",
			method: "POST",
			tmpMenu: []model.Item{
				{ItemID: 1, Name: "focaccia", Price: 5},
				{ItemID: 2, Name: "biancaneve", Price: 6},
				{ItemID: 3, Name: "margherita", Price: 6.5},
			},
			wantStCode: http.StatusMethodNotAllowed,
			wantBody:   `{"error":"method POST not allowed"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mh := NewMenuHandler(&mockMenuService{menu: tt.tmpMenu})

			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, "/menu", nil)
			r.Header.Set("Context-Type", "application/json")

			mh.GetMenu(w, r)

			assert.Equal(t, w.Code, tt.wantStCode)
			assert.JSONEq(t, tt.wantBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
