SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
cd .\write
go build main.go
cd ..\
move .\write\main .\docker\write\write
xcopy .\write\config\prod.ini .\docker\write\ /y
cd .\write\web
call npm install
call npm run build
xcopy .\dist ..\..\docker\write\dist /s /e /y
REM docker-compose build write