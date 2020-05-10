
# 简单部署样例（IBM Cloud部署，一台k8s集群）

## baaas-kubecluster
1, ibm cloud kubernetes

2, web terminal

3,cd /home/IBMid-270003757Y/.kube/config/kubeConfig277535145
```bash
luweb@k8s-terminal kubeConfig277535145 (⎈ mycluster/bqrvi4bd01rn5kmc0flg:default)$ ls
ca-hou02-mycluster.pem           kube-config-hou02-mycluster.yml
```

## File Storage: 
1, Ubuntu:

2,  NFS:
sudo apt install nfs-kernel-server nfs-common

3, NFS config 
```dotnetcli
/home/ubuntu/baas/data *(rw,sync,no_root_squash)

```

4, Restart NFS server
```dotnetcli
sudo /etc/init.d/nfs-kernel-server restart
systemctl restart nfs-server
```
## baas-kubeengine

1, 配置kube config

2, 启动kubeengine
export GOPROXY=https://goproxy.cn

cd baas-kubeengine
go run main.go


## baas-fabricengine
1, 配置baas-fabricengine/feconfig.yaml


2, 启动fabricengine
export GOPROXY=https://goproxy.cn

cd baas-fabricengine
go run main.go



## MySQL数据库
1, 安装mysql 
```
# server:
docker run -p 3306:3306 --name apimysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

# client:
sudo apt install mysql-client-core-5.7
mysql -uroot -p123456 -h127.0.0.1 -P3306
mysql -uroot -p123456 -h118.89.37.37 -P3306
```

2，导入数据：mysql.sql

##  baas-gateway 随便部署到其中一台centos
1, gwconfig.yaml  

2,
go run main.go


## baas-frontend
1, npm install


## 
* 以k8s-cluster搭建k8s集群
* k8s-master 和 baas-kubeengine 部署同一台centos
  * 将k8s master的$HOME/.kube/config文件 替换 kubeconfig/config
  * 修改配置文件 keconfig.yaml  
* nfs服务器和 baas-fabricengine 部署同一台centos
  * 设置 GOPATH 环境变量
  * 创建baas根目录
    * 复制 baas-template到其下
    * 创建nfs共享目录 baas-nfsshared 
  * 修改配置文件 feconfig.yaml  
  * nfs安装和配置
    * yum -y install nfs-utils rpcbind
    * id (查看当前用户的uid和gid)
    * vim /etc/exports (添加配置,相应修改)  
      ```
      /baas根目录/baas-nfsshared 192.168.1.0/24(rw,sync,insecure,anonuid=当前用户的uid,anongid=当前用户的gid)
      ```
    * exportfs -r (配置生效)
    * service rpcbind start &&  service nfs start (启动rpcbind、nfs服务)
  * 启动 baas-fabricengine 
* baas-gateway 随便部署到其中一台centos
  * 安装mysql 
    ```
    docker run -p 3306:3306 --name apimysql \
               -e MYSQL_ROOT_PASSWORD=123456 \
               -d mysql:5.7 
    ```
  * 通过 mysql.sql 初始化 mysql,对应修改dbconfig.yaml
  * 修改配置文件 gwconfig.yaml  
  * 运行 baas-gateway
* baas-frontend 随便部署到其中一台centos
  * npm run build:prod 打包
  * 用nginx部署，把打包生成的dist文件夹复制并重命名/usr/share/nginx/baas
  * 配置nginx.conf反向代理(相应修改baas-gateway地址)
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
  * 启动nginx
    ``` 
    sudo service nginx start
    ```
  * 访问 http://ip:8080 


