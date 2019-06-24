SRC := $(shell find ./three.js/src -type f -name "*.d.ts")

all: generate install

deps:
	npm install -g typedoc

api.json: $(SRC)
	typedoc --includeDeclarations --excludePrivate --out dist --json api.json ./three.js/src

generate: api.json
	go run . api.json

install:
	gopherjs install ../../...

clean:
	rm -f $(shell find ../.. -name "*.go" | grep -v _example | grep -v internal | grep -v static)
