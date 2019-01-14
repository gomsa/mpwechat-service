# 暂未将 Golang 集成到 docker 中
FROM alpine:latest
# 添加 ca 证书
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
RUN update-ca-certificates

RUN mkdir /app
WORKDIR /app
# 把二进制执行文件复制进容器
ADD mpwechat-service /app/mpwechat-service
# 复制配置文件
ADD config/ /app/config/
CMD ["./mpwechat-service"]