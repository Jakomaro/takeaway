package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jakomaro/takeaway/internal/model"
	"github.com/stretchr/testify/assert"
)

type mockMenuService struct {
	menu []model.Item
}

func (m *mockMenuService) GetMenu(ctx context.Context) ([]model.Item, error) {
	return m.menu, nil
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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			mh := NewMenuHandler(&mockMenuService{menu: tc.tmpMenu})

			w := httptest.NewRecorder()
			r := httptest.NewRequest(tc.method, "/menu", nil)
			r = r.WithContext(t.Context())

			r.Header.Set("Content-Type", "application/json")

			// Wrap GetMenu with middleware
			ValidateGetMethod(http.HandlerFunc(mh.GetMenu)).ServeHTTP(w, r)

			// Assert the StatusCode
			assert.Equal(t, tc.wantStCode, w.Code)

			// Assert the JSON message
			assert.JSONEq(t, tc.wantBody, strings.TrimSpace(w.Body.String()))

			// Assert the allowed method in case the method is not the same
			if tc.method != "GET" {
				assert.Equal(t, http.MethodGet, w.Header().Get("Allow"))
				return
			}

			// Assert the cache settings
			assert.Equal(t, "public, max-age=3600", w.Header().Get("Cache-Control"))
		})
	}
}
