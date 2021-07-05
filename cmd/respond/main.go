package main

import (
	"flag"
	"fmt"
	"github.com/amenasse/respond/cmd"
	"github.com/amenasse/respond/statuscode"
	"log"
	"net/http"
)

func HttpHandler(statusCode int) func(http.ResponseWriter, *http.Request) {

	description := statuscode.Description[statusCode]
	if description == "" {
		description = "Unknown"
	}

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s %s %s %s", r.Host, r.Method, r.Proto, r.URL)
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, description)
	}

}

func main() {

	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()
	code := cmd.GetStatusCode()
	path := "/"

	http.HandleFunc(path, HttpHandler(code))
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
