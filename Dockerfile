# 1. Base image
FROM golang:1.23.2-alpine

# 2. Set working directory inside container
WORKDIR /app

# 3. Copy go.mod and go.sum
COPY go.mod go.sum ./

# 4. Download dependencies
RUN go mod download

# 5. Copy the rest of the code
COPY . .

# 6. Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# 7. Expose port (masalan: 8080)
EXPOSE 8080

# 8. Start command
CMD ["./main"]

