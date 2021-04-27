.PHONY: build-api build-rpc build-all create-rpc-temp create-api-temp create-all-temp

ENV:=
SERVICE_NAME:=
API_ROOT_PATH:=./api/$(SERVICE_NAME)
RPC_ROOT_PATH:=./rpc/$(SERVICE_NAME)
FILE:=$(ENV)-$(SERVICE_NAME)
GOOS:=linux

build-api:
	env GOOS=$(GOOS) GO111MODULE=on go build -ldflags="-s -w" -o $(API_ROOT_PATH)/bin/$(FILE)-api -mod=vendor $(API_ROOT_PATH)/$(SERVICE_NAME).go \
    && mkdir -p $(API_ROOT_PATH)/doc \
    && goctl api plugin -plugin goctl-swagger="swagger -filename $(FILE).json" -api $(API_ROOT_PATH)/$(SERVICE_NAME).api -dir $(API_ROOT_PATH)/doc \
    && docker build -t $(ENV)-$(SERVICE_NAME)-api:v1 --build-arg ENV=$(ENV) --build-arg ROOT_PATH=$(API_ROOT_PATH) --build-arg SERVICE_TYPE=api --build-arg SERVICE_NAME=$(SERVICE_NAME) .

build-rpc:
	env GOOS=$(GOOS) GO111MODULE=on go build -ldflags="-s -w" -o $(RPC_ROOT_PATH)/bin/$(FILE)-rpc -mod=vendor $(RPC_ROOT_PATH)/$(SERVICE_NAME).go \
    && docker build -t $(ENV)-$(SERVICE_NAME)-rpc:v1 --build-arg ENV=$(ENV) --build-arg ROOT_PATH=$(RPC_ROOT_PATH) --build-arg SERVICE_TYPE=rpc --build-arg SERVICE_NAME=$(SERVICE_NAME) .

build-all: build-api build-rpc

create-rpc-temp:
	mkdir -p ./rpc/$(SERVICE_NAME) && cd ./rpc/$(SERVICE_NAME) \
    && goctl rpc template -o=$(SERVICE_NAME).proto

create-rpc-source:
	cd ./rpc/$(SERVICE_NAME) \
	&& goctl rpc proto -src=$(SERVICE_NAME).proto -dir=.

create-api-temp:
	mkdir -p ./api/$(SERVICE_NAME) && cd ./api/$(SERVICE_NAME) \
    && goctl api -o $(SERVICE_NAME).api

create-api-source:
	cd ./api/$(SERVICE_NAME) \
	&& goctl api go -api $(SERVICE_NAME).api -dir .

create-all-temp: create-rpc-temp create-api-temp
create-all-source: create-rpc-source create-api-source