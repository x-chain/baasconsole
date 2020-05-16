FROM alpine:latest AS production
WORKDIR /app/
ADD baas-fabricengine/baas-fabricengine /app/baas-fabricengine
ADD baas-fabricengine/feconfig.yaml /app/feconfig.yaml
EXPOSE 4991
ENTRYPOINT ["/app/baas-fabricengine"]