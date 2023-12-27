build:
	go build -o ./bin/websocket ./cmd/

run: build
	./bin/websocket
