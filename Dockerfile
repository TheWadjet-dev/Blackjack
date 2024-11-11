# Stage 1: Build the Go application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go application code
COPY main.go .

# Build the application
RUN go build -o blackjack main.go

# Stage 2: Create a lightweight image to run the application
FROM alpine:latest

# Install ca-certificates (required for Go web servers)
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /app

# Copy the built binary and HTML file from the builder stage
COPY --from=builder /app/blackjack .
COPY index.html .

# Expose port 8080 for the server
EXPOSE 8080

# Run the application
CMD ["./blackjack"]
