# Build stage
FROM --platform=linux/arm64 golang:1.21.4-bullseye AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main

# Final stage
FROM --platform=linux/arm64 gcr.io/distroless/static-debian11
COPY --from=build /app/app /app/
COPY  ./config.env /config.env
EXPOSE 8001
ENTRYPOINT ["/app/main"]