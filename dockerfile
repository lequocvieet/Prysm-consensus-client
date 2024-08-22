# Use the official Go image with the latest version as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Expose port 8080 (or whatever port your application uses)
EXPOSE 8080

# Install additional dependencies if needed (e.g., Git)
RUN apt-get update && apt-get install -y git

# Copy the Go module files to cache dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Command to run your application (you can change this as needed)
CMD ["go", "run", "main.go"]
