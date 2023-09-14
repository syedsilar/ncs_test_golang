# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code and dependencies into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the appropriate port
EXPOSE 8082 

# Command to run the application
CMD ["./main"]
