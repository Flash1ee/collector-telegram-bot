run:
	go run ./cmd/main.go -t "TOKEN" -c "./config/config.toml"

test:
	go test ./...

build:
	go build -o app.out -v ./cmd/main.go

clean:
	rm -rf *.out *.exe *.html *.csvg