# Use an official Golang runtime as a parent image
FROM golang:latest
# Set ENV variables based on the ARGs
ENV PORT=$PORT
ENV RAPIDAPI_KEY=$RAPIDAPI_KEY

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the Go app
RUN make build

# Expose port 8080 for incoming traffic
EXPOSE 8000

# Define the command to run the app when the container starts
CMD ["make", "run"]