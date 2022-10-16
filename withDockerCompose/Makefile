.DEFAULT_GOAL = help
target=dev
tag=latest

## —— 📦 The amazing Makefile 📦 —————————————————————————————————————————

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— 🐋 Docker 🐋 —————————————————————————————————————————

build: ## Up the docker environment
	target=${target} docker-compose build --progress=plain
up: ## Up the docker environment
	target=${target} docker-compose up -d
down: ## Down the docker environment
	target=${target} docker-compose down
release:
	target=prod docker-compose build --no-cache --progress=plain
run-prod:
	docker run --rm --name app-prod go-dev:19.2-prod

## —— go container —————————————————————————————————————————

enter: ## Access to container
	docker exec -u docker-user -ti go-dev /usr/bin/fish
run: ## run main.go
	docker exec -ti go-dev go run main.go
debug: ## run on debug mode, remember to listen on the IDE
	docker exec -ti go-dev /go/bin/dlv --headless --listen=:40000 --api-version=2 debug
debug-test: ## run on debug mode, remember to listen on the IDE
	docker exec -ti go-dev /go/bin/dlv --headless --listen=:40000 --api-version=2 test

