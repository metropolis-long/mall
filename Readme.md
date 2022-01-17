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

```
docker build -t [镜像的名字及标签，通常 name:tag] -f [指定要使用的Dockerfile路径]  [ContextPath]
docker build -t mall:v1 -f Dockerfile .

docker run -it [id] /bin/bash 
docker run  -p 8888:8888 -it 504a /bin/bash 

docker exec -it 243c32535da7 /bin/bash



docker images
docker ps -a 
docker kill id
docker rm id
docker rmi id


docker network create -d bridge --subnet 192.168.0.0/24 --gateway 192.168.0.1 localNet
```