# Use the official Golang image with version 1.23
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Download dependencies and tidy up the modules
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o cli-tool .

# Start a new stage from a smaller image
FROM alpine:latest

# Install necessary dependencies for running the Go binary
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/cli-tool .

# Command to run the executable
ENTRY
