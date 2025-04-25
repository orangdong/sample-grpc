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
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .

# Final stage
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Expose the gRPC port
EXPOSE 50051

# Run the server
CMD ["./server"] 