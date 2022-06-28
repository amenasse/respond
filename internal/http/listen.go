package http

import (
	"log"
	"net/http"
)

func Serve(address string, code int, body string, headers map[string]string) {
	path := "/"
	http.HandleFunc(path, Handler(code, body, headers))
	log.Fatal(http.ListenAndServe(address, nil))
}
