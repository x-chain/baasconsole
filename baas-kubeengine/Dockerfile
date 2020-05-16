FROM alpine:latest AS production
WORKDIR /app/
ADD baas-kubeengine/baas-kubeengine /app/baas-kubeengine
ADD baas-kubeengine/keconfig.yaml /app/keconfig.yaml
EXPOSE 5991
ENTRYPOINT ["/app/baas-kubeengine"]