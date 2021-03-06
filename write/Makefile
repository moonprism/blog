BINARY=main
GOOS=linux
GOARCH=amd64
WEB_DIR=./web
GO=~/code/go16/go/bin/go
DICT_TXT=~/go/pkg/mod/github.com/go-ego/gse@v0.50.3/data/dict/dictionary.txt
# for test
STOP_TOKENS_TXT=~/go/pkg/mod/github.com/go-ego/riot@v0.0.0-20201013133145-f4c30acb3704/data/dict/stop_tokens.txt

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## swag: generate api doc
swag:
	swag init

reindex:
	$(GO) run main.go reindex

data/dict.txt: $(DICT_TXT)
	mkdir -p data
	cp $(DICT_TXT) ./data/dict.txt

data/stop_tokens.txt: $(STOP_TOKENS_TXT)
	mkdir -p data
	cp $(STOP_TOKENS_TXT) ./data/stop_tokens.txt

## serve: run local dev serve
serve: data/dict.txt data/stop_tokens.txt
	$(GO) run main.go --env dev

serve-reindex: data/dict.txt data/stop_tokens.txt
	$(GO) run main.go --env dev --reindex

## serve-web: run local dev serve (web)
serve-web:
	cd $(WEB_DIR) && npm run dev

## build: build binary "main"
build: data/dict.txt data/stop_tokens.txt
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(BINARY) -v

## build-web: build frontend dist dir
build-web:
	cd $(WEB_DIR) && npm install && npm run build
