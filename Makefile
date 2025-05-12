SHELL := /bin/bash -o pipefail

APP_NAME = email-blast
PACKAGE_NAME = $(APP_NAME)
SEND_NOTIFICATION?=send_mail

.PHONY: test test-* format build

format:
	gofmt -w .

lint:
	gofmt -l .

install-ubuntu-libs:
	sudo add-apt-repository -y ppa:longsleep/golang-backports
	sudo apt update -y
	sudo apt install -y golang-go

install-tools: install-ubuntu-libs 

test: 
	export TEST_ONLY=1 && go run main.go

send-notification:
	go run main.go

