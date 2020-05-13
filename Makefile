# todo 未来会和持续集成整合
build:
	docker-compose build write

re-write:
	docker-compose stop write
	docker-compose rm write
	docker-compose build write
	docker-compose up -d

config:
	git config commit.template .git-commit-template.txt
up:
	docker-compose up -d
read-dev:
	cd read & npm run serve
write-dev:
	cd write & go run main.go
write-web-dev:
	cd write/web & npm run dev
sh-read:
	docker-compose exec read /bin/bash
sh-write:
	docker-compose exec write /bin/bash
