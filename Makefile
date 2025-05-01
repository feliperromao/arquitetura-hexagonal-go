APP_NAME=appproduct

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

bash:
	docker compose exec app bash

restart: down up

rebuild: down build up
