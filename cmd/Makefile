build:
	go build -o bin/api main.go

run: build
	GIN_MODE=release ./bin/api

clean:
	rm -r bin/

all: build