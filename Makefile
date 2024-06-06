build:
	@go build -o bin/catfact

run: build
	@ ./bin/catfact