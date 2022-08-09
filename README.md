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

Respond will bind to port 8080 on all interfaces. A different port and address can be specified:

```bash
respond -bind 127.0.0.1 -port 9090
```


To listen for HTTPS connections provide a private key and cert in PEM format using the `key` and `cert` options:

```bash
respond -key ./certs/example.com/privkey.pem -cert ./certs/example.com/cert.pem
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


All header names and values can be returned with `.RequestHeaders`

```bash

respond 200 '{{range .RequestHeaders }}{{.Name}}: {{.Value}}|{{end}}'
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
