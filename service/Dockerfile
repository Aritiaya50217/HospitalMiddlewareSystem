# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV TZ=UTC

RUN apk add --no-cache git

# cache dependency
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o service ./cmd

# ---------- Runtime stage ----------
FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/service .

EXPOSE 8080

CMD ["./service"]
