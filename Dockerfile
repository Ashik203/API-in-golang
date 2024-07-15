FROM golang:1.16.3-alpine3.13

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory to the container's working directory
COPY . .

# Set the working directory for building the Go application
WORKDIR /app/cmd/server

# Download and install dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o /app/app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/app"]
