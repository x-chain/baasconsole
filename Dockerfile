FROM golang:alpine AS builder
ENV GOPROXY http://goproxy.cn
ENV GO111MODULE on

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o kubeengine .

FROM alpine:latest AS production
WORKDIR /app/
COPY --from=builder /build/kubeengine .
EXPOSE 8080
ENTRYPOINT ["/app/kubeengine"]