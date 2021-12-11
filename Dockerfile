FROM golang:1.16

ENV GO111MODULE=on

ENV GOPROXY https://goproxy.cn,direct

WORKDIR app

COPY . .

RUN go mod tidy &&  go build  -o gin-skeleton .

EXPOSE 8081

ENTRYPOINT ["./gin-skeleton"]
