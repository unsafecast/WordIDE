
GO=go
BIN=wide

.PHONY: build clean install

build:
	$(GO) build -o $(BIN)

clean:
	go clean

install:
	cp $(BIN) /usr/bin
