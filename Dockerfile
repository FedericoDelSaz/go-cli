# Use the official Golang image as the base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only go.mod initially
COPY sre-cli-tool/go.mod ./

# Skip go.sum if it doesn't exist
RUN go mod download || true

# Download dependencies
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

# Expose the port the app runs on (if applicable)
# EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./cli-tool"]
