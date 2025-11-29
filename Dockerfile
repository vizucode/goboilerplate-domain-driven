# ============================================
# 1. STAGE BUILDER
# ============================================
FROM golang:1.25-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /app/app ./cmd/api/main.go

# ============================================
# 2. FINAL IMAGE
# ============================================
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

# copy migration folders
COPY --from=builder /app/internal/infra/migrations ./internal/infra/migrations

EXPOSE 8080

ENTRYPOINT ["./app"]
