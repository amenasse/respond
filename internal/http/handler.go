package http

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

type ResponseContext struct {
	Host          string
	StatusCode    int
	requestHeader *http.Header
	Method        string
	Proto         string
	Path          string
}

func (r ResponseContext) Description() string {
	description := StatusCodeDescription[r.StatusCode]
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

func logRequest(host string, headers http.Header, method string, protocol string, path string) {
	if host == "" {
		host = "''"
	}
	log.Printf("%s %s %s %s", host, method, protocol, path)
}

func Handler(statusCode int, responseText string, headers map[string]string) func(w http.ResponseWriter, r *http.Request) {

	context := ResponseContext{StatusCode: statusCode}
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Turn logging request headers into an option
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(string(requestDump))

		logRequest(r.Host, r.Header, r.Method, r.Proto, r.URL.String())
		for h, v := range headers {
			w.Header().Set(h, v)
		}
		w.WriteHeader(statusCode)

		context.requestHeader = &r.Header
		context.Host = r.Host
		context.Method = r.Method
		context.Proto = r.Proto
		context.Path = r.URL.String()

		t, err := template.New("response").Parse(responseText)
		err = t.Execute(w, context)
		if err != nil {
			panic(err)
		}
	}

}
