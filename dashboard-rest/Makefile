all: build

build: 
	go build -o build/micro-go ./cmd/server

run: build
	./build/micro-go

clean:
	rm -rf build

.PHONY: all build run clean
