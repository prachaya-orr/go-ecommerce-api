###################
# Docker commands
up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

clean:
	sudo rm -rf db/data


migrate:
	migrate -path [dir/migrations] -database ['url_database']

###################
# Run Main and Swagger
.PHONY: run

swagger-up:
	swag init -g main.go -o docs

test:
	go test ./...
