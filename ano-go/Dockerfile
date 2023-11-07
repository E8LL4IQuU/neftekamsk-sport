FROM golang:alpine

WORKDIR /app

RUN apk update && apk add libc-dev gcc make

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

# Install CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint.sh

ADD https://github.com/eficode/wait-for/releases/download/v2.2.4/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]

