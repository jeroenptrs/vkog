# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY *.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o vkog -ldflags="-s -w"

# Runtime stage
FROM alpine:latest

# Install dependencies
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/vkog .

# Create data directory
RUN mkdir -p /data

# Set default environment variables
# ENV VKOG_MAX_SIZE=134217728
ENV VKOG_FILE=/data/vkog.data

# Expose default port
EXPOSE 3110

# Volume for persistent storage
VOLUME ["/data"]

# Run the application
CMD ["sh", "-c", "./vkog -s ${VKOG_MAX_SIZE} -f ${VKOG_FILE}"]