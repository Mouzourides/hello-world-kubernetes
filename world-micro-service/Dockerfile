# stage 1
FROM golang:1.9.4

COPY world-micro-service.go /go/src/github.com/mouzourides/world/
WORKDIR /go/src/github.com/mouzourides/world/
RUN CGO_ENABLED=1 go build -tags netgo -a -v

# stage 2
FROM alpine:latest
COPY --from=0 /go/src/github.com/mouzourides/world/world /app/world
EXPOSE 8081
CMD ./app/world
