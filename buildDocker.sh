#!/bin/bash
# 当前目录
export BASE=$(dirname "${BASH_SOURCE[0]}") 
buildOutput=$BASE/build
# if [ ! -d "build" ];then
# mkdir build
# else
# rm -rf build/*
# fi
execDate=`date +%Y%m%d`
baasGateway="baas-gateway"
baasFabricengine="baas-fabricengine"
baasKubeengine="baas-kubeengine"

echo "编译baas-gateway,baas-fabricengine,baas-kubeengine"
docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-gateway;go build .;cd ../baas-fabricengine;go build .;cd ../baas-kubeengine;go build ."
# echo "编译baas-gateway"
# docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-gateway;go build ."
# cd $BASE/baas-gateway
# go build .
# mv baas-gateway/baas-gateway $buildOutput

tar -cvzf $buildOutput/baas-gateway-$execDate.tar.gz baas-gateway/baas-gateway baas-gateway/Dockerfile baas-gateway/*.yaml
docker build -t "baas-gateway:$execDate" -f baas-gateway/Dockerfile - < $buildOutput/baas-gateway-$execDate.tar.gz
# echo "编译baas-fabricengine"
# docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-fabricengine;go build ."
# cd $BASE/baas-fabricengine
# go build .
# mv baas-fabricengine/baas-fabricengine $buildOutput
tar -cvzf $buildOutput/baas-fabricengine-$execDate.tar.gz baas-fabricengine/baas-fabricengine baas-fabricengine/Dockerfile baas-fabricengine/*.yaml
docker build -t "baas-fabricengine:$execDate" -f baas-fabricengine/Dockerfile - < $buildOutput/baas-fabricengine-$execDate.tar.gz

# echo "编译baas-kubeengine"
# docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-kubeengine;go build ."
# cd $BASE/baas-kubeengine
# go build .
# mv baas-kubeengine/baas-kubeengine $buildOutput
tar -cvzf $buildOutput/baas-kubeengine-$execDate.tar.gz baas-kubeengine/baas-kubeengine baas-kubeengine/Dockerfile baas-kubeengine/*.yaml
docker build -t "baas-kubeengine:$execDate" -f baas-kubeengine/Dockerfile - < $buildOutput/baas-kubeengine-$execDate.tar.gz

echo "编译baas-frontend"
docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it node bash -c "cd baas-frontend;rm -rf node_modules;npm install --registry=https://registry.npm.taobao.org;npm run build:prod"
# cd $BASE/baas-frontend
# rm -rf node_modules && npm install --registry=https://registry.npm.taobao.org
# npm run build:prod
mv baas-frontend/dist $buildOutput/baas-frontend
tar -cvzf $buildOutput/baas-frontend-$execDate.tar.gz baas-frontend/Dockerfile baas-frontend/etc/baas.conf
docker build -t "baas-frontend:$execDate" - < $buildOutput/baas-frontend-$execDate.tar.gz

