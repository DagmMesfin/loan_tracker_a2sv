# Use the latest official Golang image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./
COPY main.go .

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the Go application
RUN go get
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8000

# Set the command to run the application
CMD ["./main"]