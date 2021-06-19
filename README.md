# Respond

Returns a response with the specified status code for any request

## Rationale

A http server returning specific status codes is useful for infrastructure testing and development. This is possible using a variety of multi purpose tools (netcat, socat etc..), however its not always immediately obvious what incantation is required to achieve this.


## Installation

    go get github.com/amenasse/respond

## Usage

Listen on 8080 and return a 404 status code for all requests:

```console
$ respond 404
```

Without any arguments a 200 status code is returned

```console
$ respond
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