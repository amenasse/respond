package http

import (
	"log"
	"net/http"
)

func ListenerTLS(address string, certFile string, keyFile string) func() error {
	return func() error {
		return http.ListenAndServeTLS(address, certFile, keyFile, nil)
	}
}

func Listener(address string) func() error {
	return func() error {
		return http.ListenAndServe(address, nil)
	}
}

func Serve(listener func() error, code int, body string, headers map[string]string) {
	path := "/"
	http.HandleFunc(path, Handler(code, body, headers))
	log.Fatal(listener())
}
