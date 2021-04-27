# go-zero-demo
開發步驟

1. 建立api模板 
   ex: ``` make create-api-temp SERVICE_NAME=user```
    - 會在api目錄, 建立以service name的api服務模板 user.api
    

2. 建立rpc模板 
   ex: ``` make create-rpc-temp SERVICE_NAME=user```
    - 會在rpc目錄, 建立以service name的rpc服務模板 user.proto

3. 同時建立api, rpc模板
   ex: ``` make create-all-temp SERVICE_NAME=user```
    - 會在api, rpc目錄, 建立以service name的api, rpc服務模板
    

建立模板後
1. api請編輯 xxxx.api, rpc請編輯xxxx.proto
2. 編輯好xxxx.api後, 透過temp產生api原始碼 
   ex: ```make create-api-source SERVICE_NAME=user```
    - 會產生出api layout, 開始寫程式碼
3. 編輯好xxxx.proto後, 透過temp產生rpc原始碼
   ex: ```make create-rpc-source SERVICE_NAME=user```
    - 會產生出rpc layout, 開始寫程式碼
4. 會產生出api layout, 開始寫程式碼

建立Model
```goctl model mysql datasource -url="$datasource" -table="user" -c -dir .```


建立image
```make build-all ENV=dev SERVICE_NAME=user```


#test
```curl http://localhost:8888/users/register -X POST -d '{"username": "andy","email": "andy@test.com","password": "123456"}' --header "Content-Type: application/json"```
```curl http://localhost:8888/users/login -X POST -d '{"email": "andy@test.com","password": "123456"}' --header "Content-Type: application/json"```
```curl -i -X POST http://localhost:8888/users/userinfo -H 'authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTkxODkxNzksImlhdCI6MTYxOTEwMjc3OSwidXNlcklkIjoxfQ.xJQfuPkKy7qxdQEw1itRT6ZhzIjS5AHkUdHygk-phiQ' -H 'x-user-id:1'```