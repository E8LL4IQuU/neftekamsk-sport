#!/bin/sh
# Watch for .go files change
$GOPATH/bin/CompileDaemon --build="go build -o ano-go ." --command=./ano-go
