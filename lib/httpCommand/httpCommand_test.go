package httpCommand

import (
	"net/http"
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
