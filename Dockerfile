# Start with a lightweight version of the official Golang image
FROM golang:1.21-alpine as builder

# Set necessary environment variables needed by the Go application
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /pb
WORKDIR /pb

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o pocketbase .

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory in the container
WORKDIR /pb

# Copy the pre-built binary file from the previous stage
COPY --from=builder /pb/pocketbase /pb/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]
