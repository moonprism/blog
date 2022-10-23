BINARY=linux_amd64_api
GOOS=linux
GOARCH=amd64
WEB_DIR=./web/app
GO=go

# docker build
# DICT_TXT=./mod/github.com/go-ego/gse@v0.50.3/data/dict/dictionary.txt
# for test
# STOP_TOKENS_TXT=./mod/github.com/go-ego/riot@v0.0.0-20201013133145-f4c30acb3704/data/dict/stop_tokens.txt

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## swag: generate api doc
swag:
	swag init

fmt:
	gofmt -s -w .

data/dictionary.txt:
	mkdir -p data
	wget https://raw.githubusercontent.com/go-ego/gse/v0.66.0/data/dict/dictionary.txt -O ./data/dictionary.txt

data/stop_tokens.txt:
	mkdir -p data
	wget https://raw.githubusercontent.com/go-ego/riot/master/data/dict/stop_tokens.txt -O ./data/stop_tokens.txt

## reindex: reindex
reindex: data/dictionary.txt data/stop_tokens.txt
	$(GO) run main.go se reindex

## serve: run local dev serve
serve: data/dictionary.txt data/stop_tokens.txt
	$(GO) run main.go web

## serve-web: run local dev serve (web)
serve-web:
	cd $(WEB_DIR); npm run dev

## build: build binary "main"
build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(BINARY) -v

## build-web: build frontend dist dir
build-web:
	cd $(WEB_DIR); npm install; npm run build

## docker-build-api: docker build xx
docker-build-api:
	docker-compose -f build.yml run api

build-api-before: data/dictionary.txt data/stop_tokens.txt

docker-build-web:
	docker-compose -f build.yml run web
