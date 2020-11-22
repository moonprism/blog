# TODO 未来会和持续集成整合

DOCKER=docker
DOCKER_COMPOSE=sudo docker-compose#.exe
NPM=npm

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

# Dev

## config: 配置开发环境
config:
	git config commit.template .github/.git-commit-template.txt

## up: 启动docker服务
up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

ps:
	$(DOCKER_COMPOSE) ps

logs:
	$(DOCKER_COMPOSE) logs -f

build:
	$(DOCKER_COMPOSE) build

## read-dev: 启动博客前台服务
read-dev:
	cd read && $(NPM) install && $(NPM) run serve

## read-composer-install: 博客前台 PHP vendor init
read-composer-install:
	$(DOCKER_COMPOSE) exec composer composer install -vvv

read-composer-update:
	$(DOCKER_COMPOSE) exec composer composer update

## read-npm-build: 博客前台npm run build 
read-npm-build:
	cd read && $(NPM) install && $(NPM) run build

## write-dev: 启动博客后台接口服务
write-dev:
	cd $(WRITE_DIR) && make serve

## write-web-dev: 启动博客后台web服务
write-web-dev:
	cd $(WRITE_DIR) && make serve-web

DOCKER_DIR=./docker/
WRITE_DIR=./write/
DOCKER_WRITE_DIR=$(DOCKER_DIR)write/

init:
	cd $(WRITE_DIR)config && cp test.ini prod.ini

# Write

## build-write: 博客后台容器打包编译
build-write:
	cd $(WRITE_DIR) && make build
	cd $(WRITE_DIR) && make build-web
	mv $(WRITE_DIR)main $(DOCKER_WRITE_DIR)
	mkdir -p $(DOCKER_WRITE_DIR)config && cp -r $(WRITE_DIR)config/*.ini $(DOCKER_WRITE_DIR)config
	mkdir -p $(DOCKER_WRITE_DIR)web && cp -r $(WRITE_DIR)web/dist/* $(DOCKER_WRITE_DIR)web
	$(DOCKER_COMPOSE) build write

## sh-write: 进入博客后台容器shell
sh-write:
	$(DOCKER_COMPOSE) exec write /bin/bash

## re-write: 重编译博客后台服务
re-write:
	$(DOCKER_COMPOSE) stop write
	$(DOCKER_COMPOSE) rm write
	$(DOCKER_COMPOSE) build --no-cache write
	$(DOCKER_COMPOSE) up -d

# Read

## build-read: 博客前台编译
build-read:
	cd read && npm run build

## sh-read: 进入博客前台容器shell
sh-read:
	$(DOCKER_COMPOSE) exec read /bin/bash

## sh-redis: 进入redis容器shell
sh-redis:
	$(DOCKER_COMPOSE) exec redis /bin/bash

## sh-mysql: 进入mysql容器shell
sh-mysql:
	$(DOCKER_COMPOSE) exec mysql /bin/bash

sh-composer:
	$(DOCKER_COMPOSE) exec composer /bin/bash
