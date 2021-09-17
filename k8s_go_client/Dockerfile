FROM golang:1.16-alpine AS builder

ENV CGO_ENABLE 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/demo
COPY . .
RUN  go mod tidy && go build -ldflags="-s -w" -o /app/demo cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/demo /app/demo

CMD ["./demo"]
