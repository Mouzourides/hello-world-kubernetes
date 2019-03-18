# stage 1
FROM golang:1.9.4

COPY hello-micro-service.go /go/src/github.com/mouzourides/hello/
WORKDIR /go/src/github.com/mouzourides/hello/
RUN CGO_ENABLED=1 go build -tags netgo -a -v

# stage 2
FROM alpine:latest
COPY --from=0 /go/src/github.com/mouzourides/hello/hello /app/hello
EXPOSE 8080
CMD ./app/hello
