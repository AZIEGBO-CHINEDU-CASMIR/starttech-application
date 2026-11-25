# Stage 1: Build the Go application
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Install Git and certificates
RUN apk add --no-cache git ca-certificates

# Copy Go dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o server ./cmd/api


# Stage 2: Create the lightweight production image
FROM alpine:3.21

WORKDIR /app

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Copy the compiled application
COPY --from=builder /app/server ./server

# Run as a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

# Backend application port
EXPOSE 8080

# Start the application
CMD ["./server"]