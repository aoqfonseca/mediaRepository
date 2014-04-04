setup:
	@go build .

run:
	@shell media_server

test:
	@go test ./...
