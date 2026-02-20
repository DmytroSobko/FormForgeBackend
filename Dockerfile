# =========================
# Build stage
# =========================
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server

# =========================
# Runtime stage
# =========================
FROM gcr.io/distroless/base-debian12

# IMPORTANT: match runtime working directory
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/server ./server

COPY --from=builder /app/configs ./configs

EXPOSE 8080

ENTRYPOINT ["./server"]
