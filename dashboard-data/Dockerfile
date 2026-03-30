FROM debian:bullseye-slim

WORKDIR /app

# Copy binary from local build
COPY bin/server .

EXPOSE 50051

ENTRYPOINT ["./server"]
