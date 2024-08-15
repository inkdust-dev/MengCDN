# 使用 Go 官方镜像作为基础镜像
FROM golang:1.21.10-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o cdn .

# 使用 alpine 镜像作为基础镜像
FROM alpine:latest

# 将工作目录切换为 /app
WORKDIR /app

# 将之前构建好的可执行文件拷贝到当前工作目录
COPY --from=builder /app/cdn /app/cdn

# 配置环境变量
ENV PORT=8001

# 暴露端口
EXPOSE 8001

# 增加可执行文件权限
RUN chmod +x cdn

# 配置容器启动后执行的命令
CMD ["./cdn"]
