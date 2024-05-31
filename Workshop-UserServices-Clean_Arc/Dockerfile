FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Explicitly set GOOS and GOARCH
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o main ./app


# Start from a lightweight Alpine image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=builder /app/main .

# Copy the configuration file
COPY --from=builder /app/config.env /config.env

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]

