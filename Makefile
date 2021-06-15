SERVICE_PATH := ./cmd/converter-service/main.go

OS_NAME := $(shell uname -s | tr A-Z a-z)
OS_ARH :=  $(shell arch | tr A-Z a-z)

run:
	go run $(SERVICE_PATH)

build:
	go build -o bin/app-$(OS_NAME)-$(OS_ARH) $(SERVICE_PATH)

build-all:
	GOOS=freebsd GOARCH=386 go build -o bin/app-freebsd-i386 $(SERVICE_PATH)
	GOOS=linux GOARCH=386 go build -o bin/app-linux-i386 $(SERVICE_PATH)
	GOOS=windows GOARCH=386 go build -o bin/app-windows-i386 $(SERVICE_PATH)
	GOOS=darwin GOARCH=386 go build -o bin/app-darwin-i386 $(SERVICE_PATH)
	GOOS=freebsd GOARCH=amd64 go build -o bin/app-freebsd-amd64 $(SERVICE_PATH)
	GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64 $(SERVICE_PATH)
	GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64 $(SERVICE_PATH)
	GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 $(SERVICE_PATH)