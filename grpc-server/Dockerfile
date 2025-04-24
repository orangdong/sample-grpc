# Build stage
FROM golang:1.24.1-alpine AS builder

# Install protoc and required tools
RUN apk add --no-cache protoc protobuf-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate protobuf code
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. \
    --go-grpc_opt=paths=source_relative \
    helloworld.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server

# Final stage
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Copy the proto file (optional, for reference)
COPY --from=builder /app/helloworld.proto .

# Expose the gRPC port
EXPOSE 50051

# Run the server
CMD ["./server"] 