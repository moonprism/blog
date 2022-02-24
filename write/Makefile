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

reindex:
	$(GO) run main.go se reindex

data/dict.txt: $(DICT_TXT)
	mkdir -p data
	cp $(DICT_TXT) ./data/dict.txt

data/stop_tokens.txt: $(STOP_TOKENS_TXT)
	mkdir -p data
	cp $(STOP_TOKENS_TXT) ./data/stop_tokens.txt

## serve: run local dev serve
serve: #data/dict.txt data/stop_tokens.txt
	$(GO) run main.go web

## serve-web: run local dev serve (web)
serve-web:
	cd $(WEB_DIR); npm run dev

## build: build binary "main"
build: #data/dict.txt data/stop_tokens.txt
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(BINARY) -v

## build-web: build frontend dist dir
build-web:
	cd $(WEB_DIR); npm install; npm run build

docker-build-api:
	docker-compose -f build.yml run api

build-api-after: #data/dict.txt data/stop_tokens.txt


docker-build-web:
	docker-compose -f build.yml run web
