BINARY_NAME:=http_proxy
GO_CMD:=go
build: clean
	GOOS=linux GOARCH=amd64 ${GO_CMD} build -o bin/${BINARY_NAME}-linux-64

compile:clean
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 ${GO_CMD} build -o bin/${BINARY_NAME}-linux-64

clean:
	go clean
	rm -f bin/*

run:
	go run proxy.go