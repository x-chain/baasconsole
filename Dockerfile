# FROM golang:alpine AS builder
# ENV GOPROXY http://goproxy.cn
# ENV GO111MODULE on

# WORKDIR /build

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o kubeengine .

FROM alpine:latest AS baas-fabricengine
WORKDIR /app/
ADD build/baas-fabricengine /app/baas-fabricengine
ADD build/feconfig.yaml /app/feconfig.yaml
EXPOSE 4991
ENTRYPOINT ["/app/baas-fabricengine"]

FROM alpine:latest AS baas-kubeengine
WORKDIR /app/
ADD build/baas-kubeengine /app/baas-kubeengine
ADD build/keconfig.yaml /app/keconfig.yaml
EXPOSE 5991
ENTRYPOINT ["/app/baas-kubeengine"]

FROM alpine:latest AS baas-gateway
WORKDIR /app/
ADD build/baas-gateway /app/baas-gateway
ADD build/dbconfig.yaml /app/dbconfig.yaml
ADD build/gwconfig.yaml /app/gwconfig.yaml
EXPOSE 6991
ENTRYPOINT ["/app/baas-gateway"]


FROM nginx:1.17.9 AS baas-frontend
ADD build/baas-frontend/ /usr/share/nginx/baas
ADD build/etc/baas.conf /etc/nginx/conf.d/baas.conf
EXPOSE 8080


