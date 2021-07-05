package main

import (
	"flag"
	"fmt"
	"github.com/amenasse/respond/statuscode"
	"log"
	"net/http"
	"os"
	"strconv"
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
	args := flag.Args()

	var code int = 200

	if env_var := os.Getenv("RESPONSE_STATUS"); env_var != "" {
		var err error
		if code, err = strconv.Atoi(env_var); err != nil {
			log.Fatal("Illegal status code ")
		}
	}

	if len(args) > 0 {
		if s, err := strconv.Atoi(args[0]); err == nil {
			code = s
		} else {
			log.Fatal("Illegal status code ")
		}
	}

	if code < 200 || code > 599 {
		log.Fatal("Status code out of range (should be between 200-599)")
	}

	path := "/"

	http.HandleFunc(path, HttpHandler(code))
	address := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(address, nil))
}
