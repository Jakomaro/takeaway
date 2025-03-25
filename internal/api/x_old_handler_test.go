package api

// import (
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// func TestPostOrder(t *testing.T) {

// 	validOrder := `{"order_id": 1234, "items":[{"item_id":1, "margherita", "price":8.50, "prepTime":3, "quantity":2}],"total":17.00}`
// 	missingItems := `{"order_id":1234, "total":17.00}`
// 	largePayload := strings.Repeat("a", 1024*101)

// 	type args struct {
// 		body string
// 	}

// 	tests := []struct {
// 		name           string
// 		args           args
// 		wantStatusCode int
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				body: validOrder,
// 			},
// 			wantStatusCode: 200,
// 		},
// 		{
// 			name: "missing items",
// 			args: args{
// 				body: missingItems,
// 			},
// 			wantStatusCode: 400,
// 		},
// 		{
// 			name: "large payload",
// 			args: args{
// 				body: largePayload,
// 			},
// 			wantStatusCode: 400,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			r := httptest.NewRequest("POST", "/orders", strings.NewReader(tt.args.body))
// 			r.Header.Set("Content-Type", "application/json")
// 			w := httptest.NewRecorder()

// 			PostOrder(w, r)

// 			if w.Code != tt.wantStatusCode {
// 				t.Errorf("PostOrder() = %v, want %v", w.Code, tt.wantStatusCode)
// 			}

// 		})
// 	}

// }
