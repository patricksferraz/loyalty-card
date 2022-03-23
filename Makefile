.PHONY: build

build:
	docker-compose build loyalty-card

.PHONY: status logs start stop clean

ps:
	docker-compose ps

logs:
	docker-compose logs -f loyalty-card

up:
	docker-compose up -d

start:
	docker-compose start loyalty-card

stop:
	docker-compose stop loyalty-card

down:stop
	docker-compose down -v --remove-orphans

attach:
	docker-compose exec loyalty-card bash

prune:
	docker system prune --all --volumes

.PHONY: gtest

gtest:
	go test -v -cover -coverprofile coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
