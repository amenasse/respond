package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
var StatusCodeDescription = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	103: "Early Hints",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	306: "(Unused)",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	427: "Unassigned",
	428: "Precondition Required",
	429: "Too Many Requests",
	430: "Unassigned",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	509: "Unassigned",
	510: "Not Extended",
	511: "Network Authentication Required",
}

func makeHandler(statusCode int) func(http.ResponseWriter, *http.Request) {

        description := StatusCodeDescription[statusCode]
        if description == "" {
                description = "Unknown"
        }

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(statusCode)
		fmt.Fprintln(w, description)
	}

}

func main() {

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	args := os.Args[1:]


	var code int = 200
	if len(args) > 0 {
		if s, err := strconv.Atoi(args[0]); err == nil {
			code = s
		} else {
			log.Fatal("Illegal status code ")
		}
	}

	path := "/"
	http.HandleFunc(path, makeHandler(code))
	log.Fatal(http.ListenAndServe(":"+port, nil))

}