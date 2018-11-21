package httpCommand

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPRequest_SetRequest(t *testing.T) {
	type fields struct {
		client  *http.Client
		request *http.Request
	}
	type args struct {
		url       string
		arguments []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "test1",
			fields: fields{client: &http.Client{}, request: &http.Request{}},
			args:   args{url: "http://www.google.es", arguments: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hr := &HTTPRequest{
				client:  tt.fields.client,
				request: tt.fields.request,
			}
			hr.SetRequest(tt.args.url, tt.args.arguments...)
			if hr.request == nil {
				t.Fail()
			}
		})
	}
}

//Todo add test case to check answer is cached
func TestCachedResponse(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "It works")
	}
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	//Todo Modify SetREquest to add a new parameter

	t.Fail()
}
