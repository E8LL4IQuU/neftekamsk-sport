FROM golang:alpine

WORKDIR /app

RUN apk update && apk upgrade && apk add curl

# FIXME: why not just 'COPY . .'
COPY ./go.mod go.sum ./
RUN go mod tidy && go mod download && go mod verify

# Install CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# FIXME: why move entrypoint separately
COPY . .
COPY ./entrypoint.sh /entrypoint.sh

# FIXME: do we need to chmod
RUN chmod +rx /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]
