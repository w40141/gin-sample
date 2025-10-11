# web service Dockerfile
FROM golang:1.25.2-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "cmd/api/main.go"]
