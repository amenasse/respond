package main

import (
	"flag"
	"fmt"
	"github.com/amenasse/respond/cmd"
	"github.com/amenasse/respond/statuscode"
	"log"
	"net/http"
        "os"
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
        logEnv := flag.Bool("log-env", false, "log environment variables on startup")

  flag.Usage = func() {
    fmt.Fprintf(flag.CommandLine.Output(), "" +
                  "Usage: " + os.Args[0] + " <STATUS-CODE>\n\n" +
                  "  Responds to HTTP requests with STATUS-CODE.\n" +
                  "  Without any arguments 200 OK is returned. Binds to all interfaces.\n\n")

    fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
    flag.PrintDefaults()
  }
	flag.Parse()

	code := cmd.GetStatusCode()
	path := "/"

        if *logEnv == true {
            log.Printf("====== Environment Variables (disable with log-env=false) ======")
            for _, s := range os.Environ() {
                log.Printf("%s", s)

            }
            log.Printf("===== End Environment Variables =====")
        }

	http.HandleFunc(path, HttpHandler(code))
	address := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(address, nil))
}
