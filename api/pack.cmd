rd /S/Q dist
md dist
md dist\assets
md dist\configs
md dist\plugins
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go mod tidy
go build main.go

xcopy assets dist\assets /e /y
xcopy configs dist\configs /e /y
xcopy plugins dist\plugins /e /y
copy main dist\main