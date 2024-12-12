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
RUN go build -o execution-service ./cmd/execution-service

# Runtime stage
FROM alpine:latest
WORKDIR /app

# Install Python, Go, and required libraries
RUN apk add --no-cache python3 py3-pip go

# Copy the built binary from the builder stage
COPY --from=builder /app/execution-service ./

# Ensure the binary is executable
RUN chmod +x ./execution-service

# Run the service
CMD ["./execution-service"]