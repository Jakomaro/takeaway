package api

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOrderService(t *testing.T) {

	tests := []struct {
		name string
		body string
	}{
		{
			name: "success",
			body: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/order", strings.NewReader(tc.body))
			PostOrder(w, r)

		})
	}
}
