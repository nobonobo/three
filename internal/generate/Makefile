all: generate

deps:
	npm install -g typedoc

api.json:
	typedoc --includeDeclarations --excludePrivate --out dist --json api.json ./three.js/src

generate: api.json
	go run . api.json