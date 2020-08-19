# todo 未来会和持续集成整合

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

# Dev

config:
	git config commit.template .git-commit-template.txt

up:
	docker-compose up -d

DOCKER_DIR=./docker/
WRITE_DIR=./write/
DOCKER_WRITE_DIR=$(DOCKER_DIR)write/

# Write

## build-write: 博客后台容器打包编译
build-write:
	cd $(WRITE_DIR) && make build
	cd $(WRITE_DIR) && make build-web
	mv $(WRITE_DIR)main $(DOCKER_WRITE_DIR)
	mkdir -p $(DOCKER_WRITE_DIR)config && cp -r $(WRITE_DIR)config/*.ini $(DOCKER_WRITE_DIR)config
	mkdir -p $(DOCKER_WRITE_DIR)web && cp -r $(WRITE_DIR)web/dist/* $(DOCKER_WRITE_DIR)web
	sudo docker-compose build write

## write-dev: 运行博客后台接口服务
write-dev:
	cd $(WRITE_DIR) &&　make serve

## write-web-dev: 运行博客后台前端服务
write-web-dev:
	cd $(WRITE_DIR) &&　make serve-web

## sh-write: 进入博客后台容器shell
sh-write:
	docker-compose exec write /bin/bash

## re-write: 重编译博客后台服务
re-write:
	docker-compose stop write
	docker-compose rm write
	docker-compose build write
	docker-compose up -d

# Read

## read-dev: 启动前台前端服务
read-dev:
	cd read & npm run serve

## sh-read: 进入博客前台容器shell
sh-read:
	docker-compose exec read /bin/bash

## sh-redis: 进入redis容器shell
sh-redis:
	docker-compose exec redis /bin/bash
