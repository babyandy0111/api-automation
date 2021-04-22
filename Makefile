.PHONY: build-api build-rpc clean doc build-all dev

ENV:=
SERVICE_NAME:=
API_ROOT_PATH:=./api/$(SERVICE_NAME)
RPC_ROOT_PATH:=./rpc/$(SERVICE_NAME)
FILE:=$(ENV)-$(SERVICE_NAME)
GOOS:=linux

build-api:
	env GOOS=$(GOOS) go build -ldflags="-s -w" -o $(API_ROOT_PATH)/bin/$(FILE)-api $(API_ROOT_PATH)/$(SERVICE_NAME).go \
    && mkdir -p $(API_ROOT_PATH)/doc \
    && goctl api plugin -plugin goctl-swagger="swagger -filename $(FILE).json" -api $(API_ROOT_PATH)/$(SERVICE_NAME).api -dir $(API_ROOT_PATH)/doc \
    && docker build -t $(ENV)-$(SERVICE_NAME)-api:v1 --build-arg ENV=$(ENV) --build-arg ROOT_PATH=$(API_ROOT_PATH) --build-arg SERVICE_TYPE=api --build-arg SERVICE_NAME=$(SERVICE_NAME) .

build-rpc:
	env GOOS=$(GOOS) go build -ldflags="-s -w" -o $(RPC_ROOT_PATH)/bin/$(FILE)-rpc $(RPC_ROOT_PATH)/$(SERVICE_NAME).go \
    && docker build -t $(ENV)-$(SERVICE_NAME)-rpc:v1 --build-arg ENV=$(ENV) --build-arg ROOT_PATH=$(RPC_ROOT_PATH) --build-arg SERVICE_TYPE=rpc --build-arg SERVICE_NAME=$(SERVICE_NAME) .

build-all: build-api build-rpc

dev:
	GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get github.com/tal-tech/go-zero \
    && GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get github.com/tal-tech/go-zero/rest@v1.1.6 \
    && go get github.com/tal-tech/go-zero/zrpc \
    && GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get github.com/zeromicro/goctl-swagger \
    && go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2 \
    && go get -u google.golang.org/grpc@v1.29.1 \
