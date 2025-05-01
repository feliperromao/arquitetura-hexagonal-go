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

test:
  docker compose exec app go test ./...

setup-hooks:
	cp scripts/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
restart: down up

rebuild: down build up
