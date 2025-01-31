# Start from the official Golang image as the builder stage
FROM golang:1.22-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum
RUN go mod download

# Copy the source code into container
COPY . .
COPY .env ./

# Build the application
RUN go build -o /app/bin/backend main.go

# Use a clean Debian image to run the server
FROM debian:bullseye-slim

# Install necessary runtime dependencies, if there are any
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /app/bin/backend /app/backend

# Copy the .env
COPY --from=builder /app/.env .env

# Copy the sqlitedb schemas
COPY --from=builder /app/sqlitedb sqlitedb

# make the database directory
RUN mkdir -p /app/database

# Run the server executable
CMD ["/app/backend"]

# Expose the port the server listens to
EXPOSE 8080