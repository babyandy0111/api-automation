.PHONY: build clean deploy

ENV:=dev
SERVICE_NAME:=user
SERVICE_TYPE:=api
GOOS:=linux

build:
	env GOOS=$(GOOS) go build -ldflags="-s -w" -o ./bin/$(ENV)-$(SERVICE_NAME)-$(SERVICE_TYPE) ./$(SERVICE_NAME).go && \
    docker build -t $(ENV)-$(SERVICE_NAME)-$(SERVICE_TYPE):v1 --build-arg ENV=$(ENV) --build-arg SERVICE_NAME=$(SERVICE_NAME) .

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

