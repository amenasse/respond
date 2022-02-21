FROM golang:1.16.5-alpine as build
WORKDIR /build
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD . .
RUN cd cmd/respond && go build -o /respond
# copy artifacts to a clean image
FROM alpine:3.14.0
COPY --from=build /respond /respond
ENTRYPOINT [ "/respond"]
CMD ["-port", "8080"]
