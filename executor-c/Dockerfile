FROM golang:1.24

# Install required system packages for code execution
RUN apt-get update && apt-get install -y \
    gcc g++ clang make \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY ./main.go .

# Initialize Go module (if not already done in build context)
RUN go mod init executor-c && \
    go mod tidy

# Build the Go binary
RUN go build -o executor-c .

# Expose the port used by the service
EXPOSE 5001

# Run the binary
CMD ["./executor-c"]
