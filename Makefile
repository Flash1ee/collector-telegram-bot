run:
	docker-compose down --volumes
	docker build -f ./docker/Dockerfile . --tag app
	docker-compose up --build

build:
	go build -mod=vendor -o app.out -v ./cmd/main.go

test:
	go test -mod=vendor ./...

clean:
	rm -rf *.out *.exe *.html *.csv