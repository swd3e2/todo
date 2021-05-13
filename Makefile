.PHONY: all

run:
	docker-compose up --remove-orphans --build

test:
	go test -v ./...

check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models