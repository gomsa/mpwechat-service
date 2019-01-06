# 暂未将 Golang 集成到 docker 中
FROM alpine:latest
# 添加 ca 证书
# RUN apk add --update ca-certificates && \
#     rm -rf /var/cache/apk/* /tmp/*
# RUN update-ca-certificates

RUN mkdir /app
WORKDIR /app
ADD mpwechat-service /app/mpwechat-service
CMD ["./mpwechat-service"]