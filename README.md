# Respond

Returns a response with the specified status code for any request

## Rationale

A http server returning specific status codes is useful for infrastructure testing and development. This is possible using a variety of multi purpose tools (netcat, socat etc..), however its not always immediately obvious what incantation is required to achieve this.


## Installation

    go install github.com/amenasse/respond/cmd/respond



To build a Container image:

    docker|podman build -t respond .

## Usage

Listen on 8080 and return a 404 status code for all requests:

```console
$ respond 404
```


Without any arguments a 200 status code is returned

```console
$ respond
```

The response body can be customised

```console
$ respond 200 "{{.Description}} {{.StatusCode}}\n"
```

Request headers can be returned in the response

```console
$ respond 200 '🕵: {{ .RequestHeader "User-Agent"}}'
```

Headers set multiple times can be accessed with `.RequestHeaders`

```console
$ respond 200 'Cache-Control: {{range .RequestHeaders "Cache-Control"}}{{.}} {{else}}not set{{end}}'
```

The Host request header is accessed with `.Host`

```console
$ respond 200 '{{.Host}}'
```

Respond will bind to port 8080 on all interfaces. A different port can be specified:

```console
$ respond -port 9000
```

## Limitations

Binds to all interfaces.

Support for the following may be added in the future:

 - TLS
 - HTTP/2
 - h2c
 - Websockets
