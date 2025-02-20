# Stage 1: Build the Go app
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Stage 2: Run the Go app
FROM alpine:latest

# Install CA certificates and tzdata
RUN apk --no-cache add ca-certificates tzdata

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
