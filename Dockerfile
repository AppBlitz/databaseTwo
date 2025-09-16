
# Use Go 1.25 bookworm as base image
FROM golang:1.25-bookworm AS base

# Move to working directory /build
WORKDIR /src

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod ./
# go.sum 

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o databaseBackend ./cmd/app/

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["/src/databaseBackend"]
