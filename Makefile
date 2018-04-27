db:
	docker build -t tardb . && docker run --rm -t -i -p 3301:3301 tardb

run:
	go run main.go