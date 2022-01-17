FROM golang:1.17 as go-builder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"


WORKDIR /opt/mall/


# ADD go.mod .
# ADD go.sum .
# ADD app.sh .
COPY . .
RUN go mod download \
    && go build -o ./user/rpc/userapp  user/rpc/user.go \
    && go build -o order/api/orderapp  order/api/order.go
    # && chmod +x app.sh 
# 声明服务端口
EXPOSE 8888
# 启动容器时运行的命令

# CMD ["./app.sh"]