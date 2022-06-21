package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	headers := make(map[string]string)
	handler := HttpHandler(200, "OK", headers)
	handler(recorder, req)
	res := recorder.Result()
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "OK" {
		t.Errorf("expected body 'OK' got %v", string(data))
	}

	if res.StatusCode != 200 {
		t.Errorf("expected 200 Status Code got %v", res.StatusCode)
	}
}
