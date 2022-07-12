package http

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"sort"
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

// Return  values for the given header key. Concatenate multiple values
// associated with a key into a comma seperated string
func (r ResponseContext) RequestHeader(key string) string {
	if strings.ToLower(key) == "host" {
		return r.Host
	}
	return strings.Join(r.requestHeader.Values(key), ",")
}

// RequestHeaders returns an array of type headerField so
//  header pairs can be simply ranged over in templates:
// {{range .RequestHeaders}}{{.Name}}: {{.Value}}{{end}}

type headerField struct {
	Name  string
	Value string
}

// Return all header keys and values sorted by key. Concatenate multiple values
// associated with a key into a comma seperated string
func (r ResponseContext) RequestHeaders() []headerField {

	headers := r.requestHeader.Clone()
	headers.Add("Host", r.Host)

	keys := make([]string, 0, len(headers))
	for k := range headers {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	headerFields := make([]headerField, 0)
	for _, k := range keys {
		headerFields = append(headerFields, headerField{k, strings.Join(headers[k], ",")})
	}
	return headerFields
}

type LogContext struct {
	ResponseContext
}

var LogFormat = "{{.Host}} {{.Method}} {{.Path}} {{.Proto}} {{.StatusCode}} {{.Description}} {{ range .RequestHeaders}}{{.Name}}: {{.Value}} {{end}}"

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
