setup:
	go mod init github.com/username/addingdependency
	go get github.com/google/uuid
	go mod tidy

build:
	go build -o addingdependency main.go

run: build
	./addingdependency

clean:
	rm -f addingdependency
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./addingdependency
