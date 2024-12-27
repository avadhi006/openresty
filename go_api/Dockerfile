# Use the official Golang image
FROM golang:1.21-alpine

# Set the working directory
WORKDIR /app

RUN apk --no-cache add curl


# Initialize Go module and copy the Go source code
COPY main.go .
RUN go mod init go-api                 # Initialize Go module
RUN go mod tidy                        # Fetch any required dependencies

# Build the Go application
RUN go build -o go-api

# Expose port 8086
EXPOSE 8086

# Run the Go application
CMD ["./go-api"]