FROM nginx:1.17.9
# RUN cp -rf ../../bin_x86/baas-frontend ./ && cp -rf ../../etc/ ./
ADD baas-frontend/dist /usr/share/nginx/baas
ADD baas-frontend/etc/baas.conf /etc/nginx/conf.d/baas.conf

