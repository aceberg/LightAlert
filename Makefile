mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/LightAlert && \
	go mod tidy

run:
	cd cmd/LightAlert/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/LightAlert/ && \
	CGO_ENABLED=0 go build -o ../../tmp/LightAlert .