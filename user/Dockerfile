FROM golang:1.17
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=80
WORKDIR /app
ADD . .
CMD go build .
EXPOSE 8082
ENTRYPOINT ["go","run","app.go"]