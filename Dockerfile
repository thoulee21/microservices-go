FROM golang:1.20-alpine AS builder
WORKDIR /srv/go-app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o microservice


FROM alpine
WORKDIR /srv/go-app
#COPY --from=builder /srv/go-app/other-archives ./other-archives/
COPY --from=builder /srv/go-app/microservice .

CMD ["./microservice"]