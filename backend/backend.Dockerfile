# Start from the official Golang image
FROM golang:1.23-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into container
COPY . .

# Build the application
RUN go build -o main

# Command to run executable
CMD ["/app/backend"]