package api

import (
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
		tmpMenu    []model.Item
		wantStCode int
		wantBody   string
	}{
		{
			name:       "success empty",
			tmpMenu:    []model.Item{},
			wantStCode: 200,
			wantBody:   "[]",
		},
		{
			name:       "success menu",
			tmpMenu:    []model.Item{{ItemID: 1, Name: "margherita", Price: 4.5}},
			wantStCode: 200,
			wantBody:   "[{\"item_id\":1,\"name\":\"margherita\",\"price\":4.5}]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mh := NewMenuHandler(&mockMenuService{menu: tt.tmpMenu})

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/menu", nil)
			r.Header.Set("Context-Type", "application/json")

			mh.GetMenu(w, r)

			assert.Equal(t, w.Code, tt.wantStCode)
			assert.JSONEq(t, tt.wantBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
