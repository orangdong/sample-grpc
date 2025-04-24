# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install protoc and required tools
RUN apk add --no-cache protoc

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate protobuf code
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloworld.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-client

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/grpc-client .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./grpc-client"] 