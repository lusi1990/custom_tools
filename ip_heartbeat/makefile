BINARY_NAME:=ip_heartbeat
GO_CMD:=go

.PHONY: build   clean

build: clean
	GOOS=linux GOARCH=arm ${GO_CMD} build -o bin/${BINARY_NAME}-arm main.go
	GOOS=darwin GOARCH=amd64 ${GO_CMD} build -o bin/${BINARY_NAME}-darwin-64 main.go
	GOOS=linux GOARCH=amd64 ${GO_CMD} build -o bin/${BINARY_NAME}-linux-64 main.go

clean:
	go clean
	rm -f bin/*

