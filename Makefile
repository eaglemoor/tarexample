db:
	docker build -t tardb . && docker run --rm -t -i -p 3301:3301 tardb

build-tarantool:
	docker build -t tarantool:1.10.1 ./tarantool

run:
	go run main.go