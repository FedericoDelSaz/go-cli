# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy only go.mod initially
COPY sre-cli-tool/go.mod ./

# Skip go.sum if it doesn't exist
RUN go mod download || true

# Copy the rest of the application source code
COPY sre-cli-tool .

WORKDIR /app/sre-cli-tool

# Build the Go binary
RUN go build -o top_numbers main.go

# Stage 2: Create the final lightweight image
FROM debian:bullseye-20250203-slim

# Set the working directory in the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/sre-cli-tool/top_numbers .

# Specify the command to run the CLI tool
ENTRYPOINT ["/app/top_numbers"]
