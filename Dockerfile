# 基础镜像
FROM golang:1.21.4-alpine3.17 AS builder

WORKDIR /build

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o codearena main.go

FROM alpine:3.17 AS final

# 创建工作目录
WORKDIR /app

# 复制可执行文件
COPY --from=builder /build/codearena /app/
COPY --from=builder /build/pkg/config/config.yaml /app/pkg/config/config.yaml

EXPOSE 10000

ENTRYPOINT ["/app/codearena"]