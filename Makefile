info:
	echo "Command option: 'make run' and 'make build'"
run:
	go run cmd/app/main.go
build:
	go run --output dist/server -ldflags '-s -w' cmd/app/main.go