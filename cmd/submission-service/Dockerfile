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
RUN go build -o submission-service ./cmd/submission-service

# Runtime stage
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/submission-service ./

# Ensure the binary is executable
RUN chmod +x ./submission-service

# Expose port
EXPOSE 8080

# Run the service
CMD ["./submission-service"]