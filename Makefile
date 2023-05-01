cleandb:
	docker-compose -f ./docker/local/docker-compose.yml down --remove-orphans --volumes
run:
	docker-compose -f ./docker/local/docker-compose.yml down  --remove-orphans
	docker-compose -f ./docker/local/docker-compose.yml up --build

stop:
	docker-compose -f ./docker/local/docker-compose.yml down --remove-orphans

build:
	go build -mod=vendor -o app.out -v ./cmd/main.go

test:
	go test -mod=vendor ./...


verify_migrations:
	docker-compose -f ./docker/test/docker-compose.yml up --abort-on-container-exit --build --exit-code-from migrations_test
	docker-compose -f ./docker/test/docker-compose.yml down -v --rmi "local"

clean:
	rm -rf *.out *.exe *.html *.csv