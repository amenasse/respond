package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/amenasse/respond/internal/http"
)

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

func getStatusCode(args []string) int {

	code := 200
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

func main() {
	var headers map[string]string
	headers = make(map[string]string)

	port := flag.Int("port", 8080, "port to listen on")
	cert := flag.String("cert", "", "path to TLS cert (required with tls option)")
	key := flag.String("key", "", "path to private key for the cert (required with tls option)")
	addr := flag.String("bind", "0.0.0.0", "address to bind to")
	versionFlag := flag.Bool("version", false, "display version information and exit")
	logFormat := flag.String("logformat", "", "format string for logging")
	flag.Func("header", "header to include in response", func(s string) error {
		eq := strings.IndexByte(s, ':')
		if eq == -1 {
			headers[s] = ""
		} else {
			headers[s[:eq]] = s[eq+1:]
		}
		return nil
	})

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

	tls := false

	if *key != "" || *cert != "" {
		if *key == "" || *cert == "" {
			fmt.Fprintf(os.Stderr, "Both key and cert options must be specified for TLS.\n")
			os.Exit(2)
		}
		tls = true
	}

	args := flag.Args()
	body := "{{.Description}}\n"

	if len(args) > 1 {
		body = args[1]
	}

	if *versionFlag == true {
		fmt.Printf(buildVersion(version, commit, date, builtBy))
		os.Exit(0)
	}

	if *logFormat != "" {
		http.LogFormat = *logFormat
	}
	code := getStatusCode(args)
	address := fmt.Sprintf("%s:%d", *addr, *port)
	var listener func() error

	scheme := "http"

	if tls == true {
		scheme = "https"
		listener = http.ListenerTLS(address, *cert, *key)
	} else {
		listener = http.Listener(address)
	}

	log.Printf("Starting Respond %v listening on %s://%v", version, scheme, address)
	http.Serve(listener, code, body, headers)
}
