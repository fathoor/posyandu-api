FROM golang:alpine
LABEL maintainer="mfathoor.23@gmail.com" \
      name="mfathoor/posyandu-api" \
      github="https://github.com/fathoor/posyandu-api" \
      dockerhub="https://hub.docker.com/r/mfathoor/posyandu-api"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o posyandu-api

EXPOSE ${APP_PORT}

ENTRYPOINT ["/app/posyandu-api"]
