SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build gin-api

SET GOOS=windows
 go build gin-api