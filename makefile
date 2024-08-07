build:
	go build -o bin/block 

run: build 
	./bin/docker 

test: 
	go test -v 