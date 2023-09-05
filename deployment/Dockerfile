FROM golang:latest as builder
WORKDIR /app
COPY ../ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./webapp ./cmd/main.go

FROM scratch
COPY --from=builder /app/webapp /webapp
ENTRYPOINT ['/webapp']