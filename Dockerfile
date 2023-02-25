# Build stage
FROM golang:1.20 AS build

# Set the working directory
WORKDIR /app

# Copy the source code to the container
COPY . .

# Downlaod dependencies
RUN go mod download

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o exchange


FROM alpine:latest

# Set the working directory
WORKDIR /

COPY --from=build /app/exchange /exchange

EXPOSE 4000

# Start the application
CMD ["/exchange"]
