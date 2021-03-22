.PHONY: build

build:
	go build -o build/log github.com/aurimasbachanovicius/jlog/v2/cmd/log
	go build -o build/install github.com/aurimasbachanovicius/jlog/v2/cmd/install