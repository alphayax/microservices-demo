# Start from the latest golang base image for the build stage
FROM golang:latest AS builder

# Add Maintainer Info
LABEL maintainer="Yann Ponzoni <yann.ponzoni@wescale.fr>"

# Set the Current Working Directory inside the Docker container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the Docker container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM scratch AS final

# Copy the executable from the builder stage
COPY --from=builder /app/main .

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
