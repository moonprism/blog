DOCKER=docker
DOCKER_COMPOSE=docker-compose#.exe
NPM=npm
PASSWORD=$(shell cat ./password.txt)

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

# Dev

## conf: 配置开发环境
conf:
	git config commit.template .github/.git-commit-template.txt

up:password.txt
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

ps:
	$(DOCKER_COMPOSE) ps

logs:
	$(DOCKER_COMPOSE) logs -f

## lo-read-web-dev: 本地环境启动READ web
lo-read-web-dev:
	cd read; $(NPM) install; $(NPM) run serve

## lo-read-web-build: 本地环境编译READ web
lo-read-web-build:
	cd read; $(NPM) install; $(NPM) run build

## lo-write-dev: 本地环境启动WRITE api
lo-write-dev:
	cd $(WRITE_DIR); make serve

## lo-write-web-dev: 本地环境启动WRITE web
lo-write-web-dev:
	cd $(WRITE_DIR); make serve-web

DOCKER_DIR=./docker/
WRITE_DIR=./write/

protobuf: $(wildcard *.proto)
	# protoc --go_out=plugins=grpc:write/models/protodata code.proto
	#protoc --php_out=read/app/model/protobuf --plugin=protoc-gen-grpc=/usr/games/grpc_php_plugin code.proto
	protoc --php_out=read/app/model/protobuf  code.proto

# ======= Write =======

## build-write-api: 容器编译WRITE api
build-write-api: #protobuf
	cd $(WRITE_DIR); make docker-build-api;
	$(DOCKER_COMPOSE) build write-api

## build-write-web: 容器编译WRITE web
build-write-web:
	cd $(WRITE_DIR); make docker-build-web

## admin-passwd: 设置登录用户名&密码
admin-passwd:
	$(DOCKER_COMPOSE) exec write-api /www/linux_amd64_api passwd

## se-init: 初始化查询引擎
se-init:
	$(DOCKER_COMPOSE) exec write-api /www/linux_amd64_api se reindex

## sh-write: 进入后台容器shell
sh-write:
	$(DOCKER_COMPOSE) exec write-api /bin/sh

# ======= Read =======

## build-php: 编译PHP容器后续
build-php:
	$(DOCKER_COMPOSE) exec php chown www-data:www-data /var/www/html/log
	$(DOCKER_COMPOSE) exec php composer install

## build-read: 容器编译READ
build-read:
	cd read; $(DOCKER_COMPOSE) -f build.yml run web

## sh-php: 进入PHP容器shell
sh-php:
	$(DOCKER_COMPOSE) exec php /bin/sh

## sh-redis: 进入redis容器shell
sh-redis:
	$(DOCKER_COMPOSE) exec redis redis-cli

## sh-mysql: 进入mysql容器
sh-mysql:
	$(DOCKER_COMPOSE) exec mysql mysql -u root -p$(PASSWORD)

## back-sql: 备份sql
back-sql:
	$(DOCKER_COMPOSE) exec mysql mysqldump --default-character-set=utf8mb4 -uroot -p$(PASSWORD) --databases blog --tables article article_tag code comment file tag > ./blog.sql
	cp ./blog.sql ./data/$(shell date +%Y%m%d)_blog.sql
