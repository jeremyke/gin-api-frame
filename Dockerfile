# 构建镜像使用

FROM golang as builder

WORKDIR /app

COPY . ./

# 加速！！！
ENV  GOPROXY=https://goproxy.cn,direct GO111MODULE=on
# 拉取依赖
RUN  go mod tidy
# Build
RUN  go build -ldflags '-extldflags "-static"' -o ./bin/gin-api-frame_server ./cmd/api/main.go

# 为了缩小镜像体积，做分层处理
FROM centos:7

WORKDIR /app

COPY --from=builder /app/ ./

# 启动命令，多行参数使用 `,` 隔开
# ./bin/gin-api-frame run -b ./ -e .env -p 26000 -a Local
ENTRYPOINT ["./bin/gin-api-frame_server","run","-b","/app/","-e",".env.docker","-p","26000", "-a", "DevK8S"]
