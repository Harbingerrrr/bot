# Use an official Golang image as a builder
FROM golang:1.23.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go mod and sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .
COPY .env /app/.env

# Build the Go application
RUN go build -o main .

# Use a minimal base image to run the compiled binary
FROM ubuntu:22.04

# Install any required dependencies (optional)
RUN apt-get update && apt-get install -y ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Run the application
CMD ["./main"]
