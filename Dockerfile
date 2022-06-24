FROM golang:1.17-alpine as build
WORKDIR /build
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD . .
RUN cd cmd/respond && go build -o /respond
# copy artifacts to a clean image
FROM alpine:3.16
COPY --from=build /respond /respond
ENTRYPOINT [ "/respond"]
EXPOSE 8080
CMD ["-port", "8080"]
