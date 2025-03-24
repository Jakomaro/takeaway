`Intro about TDD`


Let start with browse the menu (GET /menu). This requires the API to return a list of items. So we need to create a test for the GET menu endpoint.

We know we are going to create several endpoints in this project so let's start by creating a folder to host all of them. 
In you workspace create the `api` folder. 

In this folder we are going to create a `func TestGetMenu` where we will call our handler. Remember at the moment we don't have anything created just the go.mod and the api folder. 
a handler that return a list of items, but this is the main characteristic of TDD.

Let's create a

```go

func TestGetMenu(t *testing.T) {

	tests := []struct {
		name       string
		wantStCode int
		wantBody   string
	}{
		{
			name:       "success empty",
			wantStCode: 200,
			wantBody:   "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/menu", nil)
			r.Header.Set("Context-Type", "application/json")

			mh.GetMenu(w, r)

			assert.Equal(t, w.Code, tt.wantStCode)
			assert.JSONEq(t, tt.wantBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
```