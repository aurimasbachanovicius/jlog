.PHONY: build

build:
	go build -o build/jlog github.com/aurimasbachanovicius/jlog/v2/cmd/log
	go build -o build/jlog-install github.com/aurimasbachanovicius/jlog/v2/cmd/install

binary-to-user:
	cp build/jlog /usr/local/bin/jlog
	cp build/jlog-install /usr/local/bin/jlog-install

install: build binary-to-user