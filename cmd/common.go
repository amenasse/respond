package cmd

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetStatusCode() int {

	code := 200
	if env_var := os.Getenv("RESPONSE_STATUS"); env_var != "" {
		var err error
		if code, err = strconv.Atoi(env_var); err != nil {
			log.Fatal("Illegal status code ")
		}
	}

	args := flag.Args()
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

	return code
}

func Log(headers http.Header, method string, protocol string, path string) {

	host := headers.Get("Host")
	if host == "" {
		host = "''"
	}
	log.Printf("%s %s %s %s", host, method, protocol, path)

}
