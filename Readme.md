### run user service
```
go run ./user/rpc/user.go -f ./user/rpc/etc/user.yaml
go build -o ./user/rpc/user ./user/rpc/user.go 
./user/rpc/user -f  ./user/rpc/etc/user.yaml

```
### run order service
```
go run ./order/api/order.go -f ./order/api/etc/order.yaml
```

``` 
curl -i -X GET http://localhost:8888/api/findUserLike/get/1
curl -i -X GET http://localhost:8888/api/order/get/1

```