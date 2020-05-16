FROM alpine:latest AS production
WORKDIR /app/
ADD baas-gateway/baas-gateway /app/baas-gateway
ADD baas-gateway/dbconfig.yaml /app/dbconfig.yaml
ADD baas-gateway/gwconfig.yaml /app/gwconfig.yaml
EXPOSE 6991
ENTRYPOINT ["/app/baas-gateway"]