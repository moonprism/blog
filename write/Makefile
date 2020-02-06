serve:
	swag init
	go run main.go dev
build:
	GOOS=linux GOARCH=amd64 go build main.go
	mv ./main ../docker/write
	mv ./config/dev.ini ../docker/write