package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBody_Middleware(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		body        string
		wantStatus  int
	}{
		{
			name:        "missing content-type",
			contentType: "",
			body:        `{"valid":"json"}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "invalid content-type",
			contentType: "text/plain",
			body:        `{"valid":"json"}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "valid request",
			contentType: "application/json",
			body:        `{"valid":"json"}`,
			wantStatus:  http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))

			r.Header.Set("Content-Type", tc.contentType)

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			ValidateBody(next).ServeHTTP(w, r)
			assert.Equal(t, tc.wantStatus, w.Code)

		})
	}
}
