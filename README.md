# Respond

Return a response with the specified status code for any request

## Rationale

A http server returning specific status codes is useful for infrastructure testing and development. This is possible using a variety of multi purpose tools (netcat, socat etc..), however its not always immediately obvious what incantation is required to achieve this.


## Installation

    go install github.com/amenasse/respond/cmd/respond



To run a Container image:

    docker|podman run sysant/respond


## Usage

Listen on 8080 and return a 404 status code for all requests:

```bash
respond 404
```


Without any arguments a 200 status code is returned

```bash
respond
```
### Response Body

The response body can be customised

```bash
respond 200 '{{.Host}} {{.Method}} {{.Path}} {{.Proto}} {{.StatusCode}} {{.Description}}\n'
```

Request headers can be returned in the response

```bash
respond 200 '🕵: {{ .RequestHeader "User-Agent" }}, 👻: {{.RequestHeader "Host"}}'
```

Headers set multiple times can be accessed with `.RequestHeaders`

```bash
respond 200 'Cache-Control: {{range .RequestHeaders "Cache-Control"}}{{.}} {{else}}not set{{end}}'
```

Respond will bind to port 8080 on all interfaces. A different port can be specified:

```bash
respond -port 9000

```

### Response Headers

Reponse headers can be set with the `header` option:


```bash
respond -header 'Content-Type: application/json' \
        -header 'Last-Modified: Sun, 13 May 1984 08:52:00 GMT' \
        200 '{"Model" : "T-800", "Processor": "6502"}'
```

### Logging

Logging can be customised with the `logformat` option:

```bash
respond -logformat \
        '{"Method": "{{.Method}}","Path": "{{.Path}}", "Code": "{{.StatusCode}}"}'
```


## Limitations

Binds to all interfaces.

Support for the following may be added in the future:

 - TLS
 - HTTP/2
 - h2c
 - Websockets
