FROM debian:bullseye-slim

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

# Copy binary from local build
COPY bin/server .

EXPOSE 8080

ENTRYPOINT ["./server"]
