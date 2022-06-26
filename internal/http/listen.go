package http

import (
	"log"
	"net/http"
)

func ListenAndServe(address string, handler func(http.ResponseWriter, *http.Request)) {
	path := "/"
	http.HandleFunc(path, handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
