FROM golang:1.18.10-alpine3.17 as build

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作区
WORKDIR /go/release

# 把全部文件添加到/go/release目录
ADD . .

# 编译：把cmd/main.go编译成可执行的二进制文件，命名为app
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main main.go

# 运行：使用scratch作为基础镜像
FROM alpine as prod

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
WORKDIR /app
# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/main /app/
COPY --from=build /go/release/config.yaml /app/
RUN chmod +x /app/main
# 启动服务
CMD ["/app/main"]