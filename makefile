dev:
	go run main.go --env=development

start: build run

build:
	go build -o heli-api main.go

run:
	./heli-api