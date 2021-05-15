.PHONY:

run:
	docker-compose up --remove-orphans --build

test:
	go test --short -v ./...

test_integration:
	docker run --rm -d -p 5433:5432 --name postgres_test -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=test postgres:13.2
	go test -v ./test/
	docker stop postgres_test

check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models