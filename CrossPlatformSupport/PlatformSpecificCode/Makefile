setup:
	go mod init github.com/username/platformspecificcode
	go mod tidy

build:
	go build -o platformspecificcode

run: build
	./platformspecificcode

clean:
	rm -f platformspecificcode
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./platformspecificcode
