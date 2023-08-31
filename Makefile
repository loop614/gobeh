DOCKER_COMPOSE := docker compose
DOCKER_COMPOSE_GOBEH_BACKEND := docker compose exec gobeh_backend
DOCKER_COMPOSE_GOBEH_CONSOLE := docker compose exec gobeh_console

nuke:
	$(DOCKER_COMPOSE) down --volumes
	$(DOCKER_COMPOSE) build --no-cache
	$(DOCKER_COMPOSE) up --remove-orphans

prepare_db:
	$(DOCKER_COMPOSE_GOBEH_CONSOLE) /console/bin/console

temp:
	$(DOCKER_COMPOSE_GOBEH_BACKEND) go build -o gobeh/console/console gobeh/console/main.go
	$(DOCKER_COMPOSE_GOBEH_BACKEND) gobeh/console/console
