FROM nginx:1.17.9 AS baas

WORKDIR /app/

# baas-fabricengine
COPY build/baas-fabricengine /app/baas-fabricengine
COPY build/feconfig.yaml /app/feconfig.yaml

# baas-kubeengine
COPY build/baas-kubeengine /app/baas-kubeengine
COPY build/keconfig.yaml /app/keconfig.yaml

# baas-gateway 
COPY build/baas-gateway /app/baas-gateway
COPY build/dbconfig.yaml /app/dbconfig.yaml
COPY build/gwconfig.yaml /app/gwconfig.yaml

# baas-frontend 
COPY build/baas-frontend/ /usr/share/nginx/baas
COPY build/etc/baas.conf /etc/nginx/conf.d/baas.conf

EXPOSE 4991
EXPOSE 5991
EXPOSE 6991
EXPOSE 8080


