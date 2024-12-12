# Build stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy dependencies
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build a statically linked Linux binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o websocket-service ./cmd/websocket-service

# Runtime stage
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/websocket-service ./

# Ensure the binary is executable
RUN chmod +x ./websocket-service

# Expose port
EXPOSE 8000

# Run the service
CMD ["./websocket-service"]