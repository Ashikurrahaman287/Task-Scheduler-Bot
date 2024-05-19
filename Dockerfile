# Use the official golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o task-scheduler .

# Expose a port if your application listens on one (optional)
EXPOSE 8080

# Define the command to run your application
CMD ["./task-scheduler"]
