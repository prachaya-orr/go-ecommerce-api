###################
# Docker commands
up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

clean:
	sudo rm -rf db/data

migrate:
	migrate -path pkg/databases/migrations -database 'postgresql://postgres:postgres@localhost:4444/go_ecommerce?sslmode=disable' -verbose up

install-package:
	go get <package_name>

remove-package:
	go get <package_name>@none

install-all-dependencies:
	go get .

###################
# Run Main and Swagger
.PHONY: run

swagger-up:
	swag init -g main.go -o docs

test:
	go test ./...
