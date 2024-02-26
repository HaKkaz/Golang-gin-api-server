# Use official golang image as the base image
FROM golang:latest

# Set the working directory
WORKDIR /go/src/app

# Copy all files to the container
COPY . .

# Install dependencies
RUN go mod download

# Construct go app binary
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

