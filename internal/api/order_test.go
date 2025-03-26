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
	}{
		{
			name:       "success",
			body:       "",
			wantStCode: http.StatusCreated,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/order", strings.NewReader(tc.body))
			PostOrder(w, r)

			assert.Equal(t, tc.wantStCode, w.Code)

		})
	}
}
