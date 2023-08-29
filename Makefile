DOCKER_COMPOSE := docker compose
DOCKER_COMPOSE_GOBEH_BACKEND := docker compose exec gobeh_backend

docker_rebuild:
	$(DOCKER_COMPOSE) build --no-cache
	$(DOCKER_COMPOSE) up --remove-orphans

temp:
	$(DOCKER_COMPOSE_GOBEH_BACKEND) cat core/behhandler.go

go_start:
	$(DOCKER_COMPOSE_GOBEH_BACKEND) /usr/local/bin/main
