setup:
	go mod init github.com/username/inference
	go get golang.org/x/exp/constraints
	go mod tidy

build:
	go build -o inference main.go

run:
	go run main.go

clean:
	rm -f inference
	rm -f go.mod
	rm -f go.sum
