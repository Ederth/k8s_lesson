FROM golang:1.16 AS builder

WORKDIR /src

ENV GOPROXY=https://goproxy.cn

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . /src
RUN  mkdir -p bin/ && go build -o ./bin/server ./...

FROM debian:stable-slim

RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list \
        && apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

ENV TZ=Asia/Shanghai

EXPOSE 80
CMD ["./server"]
