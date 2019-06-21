all: cli srv

cli:
	mkdir -p bin
	go build -o bin/enghack ./client

srv:
	mkdir -p bin
	go build -o bin/enghack-server ./server