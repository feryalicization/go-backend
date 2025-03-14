
FROM golang:1.18-alpine as builder

# Set the working directory for the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy all application files into the container
COPY . .

# Verify Go version (this will print the Go version in logs)
RUN go version
