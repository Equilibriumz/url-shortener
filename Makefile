.PHONY: dev

BUILD_NAME := url-shortener

dev:
	go get .
	go build -o=./bin/${BUILD_NAME}