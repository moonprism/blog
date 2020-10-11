BINARY=main
GOOS=linux
GOARCH=amd64
WEB_DIR="./web"

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## swag: generate api doc
swag:
	swag init

## serve: run local dev serve
serve:
	go run main.go --env dev

## serve-web: run local dev serve (web)
serve-web:
	cd $(WEB_DIR) && npm run dev

## build: build binary "main"
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY) -v

## build-web: build frontend dist dir
build-web:
	cd $(WEB_DIR) && npm install && npm run build
