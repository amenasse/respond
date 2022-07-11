package http

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
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

// Return all headers and values. If multiple values are associated with a key return as a comma seperated string
func (r ResponseContext) RequestHeadersAll() map[string]string {

	h := make(map[string]string)
	for k, v := range *r.requestHeader {
		h[k] = strings.Join(v, ",")
	}
	return h
}

type LogContext struct {
	ResponseContext
}

var LogFormat = "{{.Host}} {{.Method}} {{.Path}} {{.Proto}} {{.StatusCode}} {{.Description}} {{ range $key,$value := .RequestHeadersAll}}{{$key}}: {{$value}} {{end}}"

func renderRequestLog(context LogContext) string {

	if context.Host == "" {
		context.Host = "''"
	}

	t, err := template.New("log").Parse(LogFormat)
	var buf bytes.Buffer
	err = t.Execute(&buf, context)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func Handler(statusCode int, responseText string, headers map[string]string) func(w http.ResponseWriter, r *http.Request) {

	context := ResponseContext{StatusCode: statusCode}
	return func(w http.ResponseWriter, r *http.Request) {

		for h, v := range headers {
			w.Header().Set(h, v)
		}
		w.WriteHeader(statusCode)

		context.requestHeader = &r.Header
		context.Host = r.Host
		context.Method = r.Method
		context.Proto = r.Proto
		context.Path = r.URL.String()

		logContext := LogContext{ResponseContext: context}
		log.Printf(renderRequestLog(logContext))

		t, err := template.New("response").Parse(responseText)
		err = t.Execute(w, context)
		if err != nil {
			panic(err)
		}
	}

}
