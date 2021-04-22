# go-zero-demo

產生模板：```goctl api -o user.api```

編輯模板<將req,res,route定義好>

透過模板產生layout ```goctl api go -api user.api -dir .```

編輯 etc/user-api.yaml

產生model ```goctl model mysql datasource -url="$datasource" -table="user" -c -dir .```

編輯 internal/handler/*.go

編輯主邏輯 internal/logic/*.go 

go run xxx.go

make build-all ENV=dev SERVICE_NAME=user


curl http://localhost:8888/users/register -X POST -d '{"username": "andy","email": "andy@test.com","password": "123456"}' --header "Content-Type: application/json"
