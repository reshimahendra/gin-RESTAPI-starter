info:
	echo "Command option: 'make run' and 'make build'"
run:
	go run cmd/app/main.go
build:
	go run --output dist/server -ldflags '-s -w' cmd/app/main.go
pq-start:
	# starting postgresql service (if not started)
	sudo systemctl start postgresql
pq-stop:
	# stoping postgresql service
	sudo systemctl stop postgresql
pq-stat:
	# show postgresql service status
	sudo systemctl status postgresql
