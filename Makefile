.PHONY: build binary-to-user install

build:
	go build -o build/jlog github.com/aurimasbachanovicius/jlog/v2/cmd/log

binary-to-user:
	cp build/jlog /usr/local/bin/jlog

install: build binary-to-user