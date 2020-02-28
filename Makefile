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