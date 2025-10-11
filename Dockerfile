FROM golang:1.25.2-trixie AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -buildvcs=false -trimpath -ldflags '-w -s' -o myexe ./cmd/api/main.go

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/myexe /app/myexe

ENTRYPOINT ["/app/myexe"]
