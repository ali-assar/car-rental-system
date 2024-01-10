# Use the official Golang image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy only the go module files to leverage Docker caching
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the local package files to the container
COPY . .

# Download and install any required dependencies
RUN go get

# Build the application with -mod=readonly for consistency
RUN go build -o main -mod=readonly .

# Expose the port your application listens on
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
