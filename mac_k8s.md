# 简单部署样例（IBM Cloud 部署，一台 k8s 集群）

## baaas-kubecluster

1, ibm cloud kubernetes

2, web terminal

3,cd /home/IBMid-270003757Y/.kube/config/kubeConfig277535145

```bash
luweb@k8s-terminal kubeConfig277535145 (⎈ mycluster/bqrvi4bd01rn5kmc0flg:default)$ ls
ca-hou02-mycluster.pem           kube-config-hou02-mycluster.yml
```

## File Storage:

1, Mac:

2, NFS Server:

4, Restart NFS server

5, Mac Client

mount -t nfs -o rw 118.89.37.37:/Users/blockchain/baas/data /Users/blockchain/baas/data

docker run -d --name nfs --privileged -v /Users/blockchain/data:/nfsshare -e SHARED_DIRECTORY=/nfsshare itsthenetwork/nfs-server-alpine:latest

docker run -d --name nfs --restart=always --net=host --privileged -v /Users/blockchain/data:/nfsshare -e SHARED_DIRECTORY=/nfsshare itsthenetwork/nfs-server-alpine:latest

docker run -d --name nfs --restart=always -p2049:2049 --privileged -v /Users/blockchain/data:/nfsshare -e SHARED_DIRECTORY=/nfsshare itsthenetwork/nfs-server-alpine:latest

## baas-kubeengine

1, 配置 kube config

2, 启动 kubeengine
export GOPROXY=https://goproxy.cn

cd baas-kubeengine
go run main.go

## baas-fabricengine

1, 配置 baas-fabricengine/feconfig.yaml

2, 启动 fabricengine
export GOPROXY=https://goproxy.cn

cd baas-fabricengine
go run main.go

## MySQL 数据库

1, 安装 mysql

```
# server:
docker run -p 3306:3306 --name apimysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

# client:
sudo apt install mysql-client-core-5.7
mysql -uroot -p123456 -h127.0.0.1 -P3306
mysql -uroot -p123456 -h118.89.37.37 -P3306
```

2，导入数据：mysql.sql

## baas-gateway 随便部署到其中一台 centos

1, gwconfig.yaml

2,
go run main.go

## baas-frontend

1, npm install

##

- 以 k8s-cluster 搭建 k8s 集群
- k8s-master 和 baas-kubeengine 部署同一台 centos
  - 将 k8s master 的\$HOME/.kube/config 文件 替换 kubeconfig/config
  - 修改配置文件 keconfig.yaml
- nfs 服务器和 baas-fabricengine 部署同一台 centos
  - 设置 GOPATH 环境变量
  - 创建 baas 根目录
    - 复制 baas-template 到其下
    - 创建 nfs 共享目录 baas-nfsshared
  - 修改配置文件 feconfig.yaml
  - nfs 安装和配置
    - yum -y install nfs-utils rpcbind
    - id (查看当前用户的 uid 和 gid)
    - vim /etc/exports (添加配置,相应修改)
      ```
      /baas根目录/baas-nfsshared 192.168.1.0/24(rw,sync,insecure,anonuid=当前用户的uid,anongid=当前用户的gid)
      /data/baas/manager 10.2.1.0/24(rw,sync,insecure,anonuid=0,anongid=0)
      ```
    - exportfs -r (配置生效)
    - service rpcbind start && service nfs start (启动 rpcbind、nfs 服务)
  - 启动 baas-fabricengine
- baas-gateway 随便部署到其中一台 centos
  - 安装 mysql
    ```
    docker run -p 3306:3306 --name apimysql \
               -e MYSQL_ROOT_PASSWORD=123456 \
               -d mysql:5.7
    ```
  - 通过 mysql.sql 初始化 mysql,对应修改 dbconfig.yaml
  - 修改配置文件 gwconfig.yaml
  - 运行 baas-gateway
- baas-frontend 随便部署到其中一台 centos

  - npm run build:prod 打包
  - 用 nginx 部署，把打包生成的 dist 文件夹复制并重命名/usr/share/nginx/baas
  - 配置 nginx.conf 反向代理(相应修改 baas-gateway 地址)

    ```
    user www-data;
    worker_processes auto;
    pid /run/nginx.pid;

    events {
    	worker_connections 768;
    	# multi_accept on;
    }

    http {
        include       mime.types;
        default_type  application/octet-stream;

        log_format  logformat  '$remote_addr - $remote_user [$time_local] "$request" '
                          '$status $body_bytes_sent "$http_referer" '
                          '"$http_user_agent" "$http_x_forwarded_for" '
                          '"[$request_time]" "[$upstream_response_time]" '
                          '"[$connection]" "[$connection_requests]" '
                          '"$http_imei" "$http_mobile" "$http_type" "$http_key" "$cookie_sfpay_jsessionid"';
        access_log  /var/log/nginx/access.log logformat;
    ```


        sendfile        on;
        #tcp_nopush     on;
        underscores_in_headers on;

        keepalive_timeout  65;
        proxy_connect_timeout 120;
        proxy_read_timeout 120;
        proxy_send_timeout 60;
        proxy_buffer_size 16k;
        proxy_buffers 4 64k;
        proxy_busy_buffers_size 128k;
        proxy_temp_file_write_size 128k;
        proxy_temp_path /tmp/temp_dir;
        proxy_cache_path /tmp/cache levels=1:2 keys_zone=cache_one:200m inactive=1d max_size=30g;

        client_header_buffer_size 12k;
        open_file_cache max=204800 inactive=65s;
        open_file_cache_valid 30s;
        open_file_cache_min_uses 1;



        gzip  on;
        gzip_types       text/plain application/x-javascript text/css application/xml text/javascript application/x-httpd-php image/jpeg image/gif image/png image/jpg;
        # baas-gateway地址
        upstream baasapi {
            server 127.0.0.1:6991;
        }


        # HTTP server
        #
        server {
            listen       8080;
            server_name  baasadmin;

            location /nginx_status {
                    stub_status on;
                    access_log off;
            }
            location /api/{
                proxy_pass  http://baasapi/api/;
                proxy_set_header  X-Real-IP  $remote_addr;
                proxy_set_header Host $host;

            }
            location /dev-api/{
                proxy_pass  http://baasapi/api/;
                proxy_set_header  X-Real-IP  $remote_addr;
                proxy_set_header Host $host;

            }
            location /stage-api/{
                proxy_pass  http://baasapi/api/;
                proxy_set_header  X-Real-IP  $remote_addr;
                proxy_set_header Host $host;

            }

            location / {
                root   baas;
                index  index.html index.htm;
            }

            location ~ ^/favicon\.ico$ {
                root   baas;
            }

        }
    }
    ```

- 启动 nginx
  ```
  sudo service nginx start
  ```
- 访问 http://ip:8080
