FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o posyandu-api

EXPOSE ${APP_PORT}

ENTRYPOINT ["/app/posyandu-api"]
