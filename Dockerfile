# syntax=docker/dockerfile:1
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o app ./cmd/main.go

# Command to run the executable
CMD ["./app"]