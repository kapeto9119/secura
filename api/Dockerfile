# Build stage
FROM golang:1.20-alpine AS builder

# Install git for fetching dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o secura-api ./cmd/server

# Use a minimal alpine image for the final container
FROM alpine:3.18

# Add CA certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/secura-api .

# Expose port
EXPOSE 8080

# Start the application
CMD ["./secura-api"] 