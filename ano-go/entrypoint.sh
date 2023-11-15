echo "Pinging database"
wait-for "${DATABASE_CONTAINER_NAME}:${DATABASE_PORT}" -- "$@"

# Watch for .go files change
$GOPATH/bin/CompileDaemon --build="go build -o ano-go ." --command=./ano-go
