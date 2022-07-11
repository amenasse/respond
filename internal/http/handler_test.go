package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	type test struct {
		statusCode      int
		bodyTemplate    string
		responseBody    string
		responseHeaders map[string]string
		requestHeaders  map[string][]string
	}

	tests := []test{
		{
			statusCode:   200,
			bodyTemplate: "OK",
			responseBody: "OK",
		},
		{
			statusCode:   200,
			bodyTemplate: "{{.StatusCode}} {{.Description}}",
			responseBody: "200 OK",
		},
		{
			statusCode:   444,
			bodyTemplate: "{{.Description}}",
			responseBody: "Unknown",
		},
		{
			statusCode:     444,
			bodyTemplate:   "{{.RequestHeader \"Host\"}}",
			responseBody:   "example.com",
			requestHeaders: map[string][]string{"Host": {"example.com"}},
		},
		{
			statusCode:     200,
			bodyTemplate:   "{{.RequestHeader \"User-Agent\"}}",
			responseBody:   "Mosaic/0.9",
			requestHeaders: map[string][]string{"User-Agent": {"Mosaic/0.9"}},
		},

		{
			statusCode:   200,
			bodyTemplate: "|{{range .RequestHeaders \"Cache-Control\"}}{{.}}|{{end}}",
			responseBody: "|max-age=0|private|",
			requestHeaders: map[string][]string{
				"Cache-Control": {"max-age=0", "private"},
			},
		},
		{
			statusCode:   200,
			bodyTemplate: "OK",
			responseBody: "OK",
			responseHeaders: map[string]string{
				"Content-Type":  "application/json",
				"Last-Modified": "Sun, 13 May 1984 08:52:00 GMT",
			},
		},
	}

	for _, tc := range tests {

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()
		for k, v := range tc.requestHeaders {
			for _, h := range v {
				req.Header.Add(k, h)
			}
		}
		handler := Handler(
			tc.statusCode,
			tc.bodyTemplate,
			tc.responseHeaders,
		)
		handler(recorder, req)
		res := recorder.Result()
		data, _ := ioutil.ReadAll(res.Body)
		if string(data) != tc.responseBody {
			t.Errorf("expected body %v got %v", tc.responseBody, string(data))
		}

		for k, v := range tc.responseHeaders {
			hv := res.Header.Get(k)
			if res.Header.Get(k) != v {
				t.Errorf("expected header %v: %v got %v", k, v, hv)
			}
		}
	}
}
