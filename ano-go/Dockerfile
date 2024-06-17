# Step 1: Build go application
FROM golang AS build

WORKDIR /app

RUN apt update && apt upgrade -y

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/ano-go .

# Step 2: Run go application
FROM debian AS prod-stage

WORKDIR /app

RUN apt update && apt upgrade -y && \
    apt install -y curl

COPY --from=build /app/ano-go .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

CMD ["./ano-go"]