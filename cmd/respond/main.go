package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/amenasse/respond/cmd"
	"github.com/amenasse/respond/statuscode"
)

type ResponseContext struct {
	Host          string
	StatusCode    int
	requestHeader *http.Header
}

func (r ResponseContext) Description() string {
	description := statuscode.Description[r.StatusCode]
	if description == "" {
		description = "Unknown"
	}
	return description
}

// Simplify referencing request headers first value  in the response template.
func (r ResponseContext) RequestHeader(key string) string {
	if strings.ToLower(key) == "host" {
		return r.Host
	}
	return r.requestHeader.Get(key)
}

// Simplify referencing request header values in the response template.
func (r ResponseContext) RequestHeaders(key string) []string {
	return r.requestHeader.Values(key)
}

func HttpHandler(statusCode int, responseText string) func(w http.ResponseWriter, r *http.Request) {

	context := ResponseContext{StatusCode: statusCode}
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Turn logging request headers into an option
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(string(requestDump))

		cmd.Log(r.Header, r.Method, r.Proto, r.URL.String())
		w.WriteHeader(statusCode)

		context.requestHeader = &r.Header
		context.Host = r.Host
		t, err := template.New("response").Parse(responseText)
		err = t.Execute(w, context)
		if err != nil {
			panic(err)
		}
	}

}

// buildVersion taken from goreleaser https://github.com/goreleaser/goreleaser
func buildVersion(version, commit, date, builtBy string) string {
	const website = "\n\nhttps://github.com/amenasse/respond"
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	result = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s", result, runtime.GOOS, runtime.GOARCH)
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
	}
	return result + website + "\n"
}

// Build related variables should be set by ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {

	port := flag.Int("port", 8080, "port to listen on")
	versionFlag := flag.Bool("version", false, "display version information and exit")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), ""+
			"Usage: "+os.Args[0]+" <STATUS-CODE> [RESPONSE-BODY]\n\n"+
			"  Responds to HTTP requests with STATUS-CODE.\n"+
			"  Without any arguments 200 OK is returned. Binds to all interfaces.\n"+
			"  RESPONSE-BODY should be a Go template\n\n")

		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()

	body := "{{.Description}}\n"
	if len(args) > 1 {
		body = args[1]
	}

	if *versionFlag == true {
		fmt.Printf(buildVersion(version, commit, date, builtBy))
		os.Exit(0)
	}

	code := cmd.GetStatusCode()
	path := "/"

	http.HandleFunc(path, HttpHandler(code, body))
	address := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(address, nil))
}
