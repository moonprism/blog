DOCKER=docker
DOCKER_COMPOSE=docker-compose#.exe
NPM=npm
PASSWORD=$(shell cat ./password.txt)

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

# Dev

## config: 配置开发环境
config:
	git config commit.template .github/.git-commit-template.txt

## up: 启动docker服务
up:password.txt
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

ps:
	$(DOCKER_COMPOSE) ps

logs:
	$(DOCKER_COMPOSE) logs -f

build:
	$(DOCKER_COMPOSE) build

## read-dev: 启动前台服务
read-dev:
	cd read; $(NPM) install; $(NPM) run serve

## read-composer-install: 前台 PHP vendor init
read-composer-install:
	$(DOCKER_COMPOSE) exec composer composer install -vvv

read-composer-update:
	$(DOCKER_COMPOSE) exec composer composer update

## read-npm-build: 前台npm run build 
read-npm-build:
	cd read; $(NPM) install; $(NPM) run build

## write-dev: 启动后台接口服务
write-dev:
	cd $(WRITE_DIR); make serve

## write-web-dev: 启动后台web服务
write-web-dev:
	cd $(WRITE_DIR); make serve-web

DOCKER_DIR=./docker/
WRITE_DIR=./write/

protobuf: $(wildcard *.proto)
	protoc --go_out=plugins=grpc:write/protodata code.proto
	#protoc --php_out=read/app/model/protobuf --plugin=protoc-gen-grpc=/usr/games/grpc_php_plugin code.proto
	protoc --php_out=read/app/model/protobuf  code.proto

# ======= Write =======

## build-write-api: 后台api容器打包编译
build-write-api: #protobuf
	cd $(WRITE_DIR); make docker-build-api; make build-api-after
	$(DOCKER_COMPOSE) build write-api

## build-write-web: 后台web编译
build-write-web:
	cd $(WRITE_DIR); make docker-build-web

## admin-passwd: 设置后台登录用户名&密码
admin-passwd:
	$(DOCKER_COMPOSE) exec write-api /www/linux_amd64_api passwd

## meili-init: 从数据库初始化查询引擎
meili-init:
	$(DOCKER_COMPOSE) exec write-api /www/linux_amd64_api se reindex

## sh-write: 进入后台容器shell
sh-write:
	$(DOCKER_COMPOSE) exec write-api /bin/sh

# ======= Read =======

# build-php: 使用当前系统代理编译php镜像
# build-php:
# 	cd $(DOCKER_DIR)php; docker build --network host .

build-php-after:
	$(DOCKER_COMPOSE) up -d ; $(DOCKER_COMPOSE) exec php composer install

## build-read: 前台编译
build-read:
	cd read; $(DOCKER_COMPOSE) -f build.yml run web

## sh-php: 进入前台容器shell
sh-php:
	$(DOCKER_COMPOSE) exec php /bin/sh

## sh-redis: 进入redis容器shell
sh-redis:
	$(DOCKER_COMPOSE) exec redis redis-cli

## sh-mysql: 进入mysql容器shell
sh-mysql:
	$(DOCKER_COMPOSE) exec mysql mysql -u root -p$(PASSWORD)

## back-sql: 备份sql
back-sql:
	$(DOCKER_COMPOSE) exec mysql mysqldump -uroot -p$(PASSWORD) --databases blog --tables article article_tag code comment file tag > ./data/$(shell date +%Y%m%d)_blog.sql
