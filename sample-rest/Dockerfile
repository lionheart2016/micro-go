FROM debian:bullseye-slim

WORKDIR /app

# Copy binary from local build
COPY bin/server .

EXPOSE 8080

ENTRYPOINT ["./server"]
