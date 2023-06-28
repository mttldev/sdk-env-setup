@echo off
echo Building...
go build setup.go sdk_urls.go web_url.go netutil.go ziputil.go fileutil.go
echo Done.
