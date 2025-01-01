setup:
	go mod init github.com/username/usinggosqlite3
	go get github.com/mattn/go-sqlite3
	go mod tidy

build:
	go build -o usinggosqlite3 main.go

run: build
	./usinggosqlite3

clean:
	rm -f usinggosqlite3
	rm -f go.mod
	rm -f go.sum
	rm -f test.db

complete: clean setup build
	./usinggosqlite3
