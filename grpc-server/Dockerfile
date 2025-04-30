# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Copy go mod files first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy only necessary source code for the build
COPY helloworld ./helloworld
COPY main.go .

# Build the application
# Using '.' builds the package in the current directory
# Explicitly set OS and Architecture for cross-compilation if needed
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server .

# Final stage
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Expose the gRPC port
EXPOSE 50051
# Expose the HTTP health check port
EXPOSE 8080

# Run the server
CMD ["./server"] 