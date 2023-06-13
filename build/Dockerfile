FROM alpine:3.14
ADD app /bin/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk -Uuv add --no-cache ca-certificates tini tzdata && \
  ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  mkdir /wx-proxy
WORKDIR /wx-proxy
ADD . .
ENTRYPOINT ["/sbin/tini","--", "app"]