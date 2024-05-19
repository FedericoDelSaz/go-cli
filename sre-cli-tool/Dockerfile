# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy only go.mod initially
COPY go.mod ./

# Skip go.sum if it doesn't exist
RUN go mod download || true

# Copy the rest of the application source code
COPY . .

# Build the Go binary
RUN go build -o cli-tool main.go

# Stage 2: Create the final lightweight image
FROM debian:bullseye-slim

# Set the working directory in the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/cli-tool .

# Specify the command to run the CLI tool
ENTRYPOINT ["/app/cli-tool"]
