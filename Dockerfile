FROM golang:1.17 as go-builder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"


WORKDIR /opt/mall/


ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/hello/etc /app/etc
RUN go build -ldflags="-s -w" -o ./user/rpc/app  user/rpc/user.go 
# 声明服务端口
EXPOSE 8080
# 启动容器时运行的命令
 CMD ["./user/rpc/app", "-f", "user/rpc/etc/user.yaml"]