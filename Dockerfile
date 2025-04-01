# --- Build Stage ---
    FROM golang:1.23.4 AS builder
    WORKDIR /app
    
    # Copy Go modules and install dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # Build the Go application
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
    
    # --- Final Stage ---
    FROM alpine:latest
    WORKDIR /root/
    
    # Install necessary dependencies
    RUN apk --no-cache add ca-certificates
    
    # Copy the compiled binary from the builder stage
    COPY --from=builder /app/main .
    COPY --from=builder /app/.env .
    
    # Ensure the binary is executable
    RUN chmod +x ./main
    
    # Set environment variables
    ENV GIN_MODE=release
    
    # Run the application
    CMD ["./main"]    