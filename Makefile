run:
	docker-compose down --volumes
	docker build -f ./docker/Dockerfile . --tag app
	docker-compose up --build

build:
	go build -o app.out -v ./cmd/main.go

test:
	go test ./...

clean:
	rm -rf *.out *.exe *.html *.csvg