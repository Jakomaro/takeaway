package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderService(t *testing.T) {

	tests := []struct {
		name       string
		body       string
		wantStCode int
		wantBody   string
	}{
		{
			name:       "valid order",
			body:       `{"order_id": 1234, "items":[{"item_id":1, "name":"margherita", "price":8.50}],"total":17.00}`,
			wantStCode: http.StatusCreated,
			wantBody:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/order", strings.NewReader(tc.body))

			r.Header.Set("Contest-Type", "application/json")

			ValidateBody(http.HandlerFunc(PostOrder)).ServeHTTP(w, r)
			assert.Equal(t, tc.wantStCode, w.Code)

			assert.Equal(t, tc.wantBody, w.Body.String())

		})
	}
}
