all: cli srv

cli:
	go build -o enghack ./client

srv:
	go build -o enghack-server ./server