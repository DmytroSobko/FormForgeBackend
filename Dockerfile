# =========================
# Build stage
# =========================
FROM golang:1.22-alpine AS builder

# Install git (needed for go modules)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Copy go.sum only if it exists
COPY go.sum ./
RUN go mod download || true
# Copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server


# =========================
# Runtime stage
# =========================
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/server /server

# Expose port
EXPOSE 8080

# Run the server
ENTRYPOINT ["/server"]