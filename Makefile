.PHONY: build

build:
	go build -o build/jlog github.com/aurimasbachanovicius/jlog/v2/cmd/log
	go build -o build/jlog-install github.com/aurimasbachanovicius/jlog/v2/cmd/install