FROM golang:1.20-alpine as vendor
RUN apk --no-cache add ca-certificates

WORKDIR /build
COPY vendor .

FROM vendor as builder

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o app -v ./cmd/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/app /app
COPY --from=builder /build/config/config.toml /config.toml

ENTRYPOINT ["/app", "-c", "/config.toml"]