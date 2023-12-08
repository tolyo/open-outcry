# Build stage
FROM golang:1.21 AS build

# Set the working directory
WORKDIR /temp

# Copy the source code to the container
COPY . .

# Downlaod dependencies
RUN go mod download

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app


FROM alpine:latest

# Set the working directory
WORKDIR /

COPY --from=build /temp/app /app

EXPOSE 4000

# Start the application
CMD ["/app"]
