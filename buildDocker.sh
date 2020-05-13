#!/bin/bash
# 当前目录
export BASE=$(dirname "${BASH_SOURCE[0]}") 
if [ ! -d "bin" ];then
mkdir bin
else
rm -rf bin/*
fi
echo "编译baas-gateway"
docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-gateway;go build ."
# cd $BASE/baas-gateway
# go build .
mv baas-gateway $BASE/bin
echo "编译baas-fabricengine"
docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-fabricengine;go build ."
# cd $BASE/baas-fabricengine
# go build .
mv baas-fabricengine $BASE/bin
echo "编译baas-kubeengine"
docker run -w /build/ -e GO111MODULE=on -eGOCACHE=on -eGOPROXY=https://goproxy.io -v`pwd`:/build -it hyperledger/fabric-tools:1.4.1 bash -c "cd baas-kubeengine;go build ."
# cd $BASE/baas-kubeengine
# go build .
mv baas-kubeengine $BASE/bin
echo "编译baas-frontend"
cd $BASE/baas-frontend
rm -rf node_modules && npm install --registry=https://registry.npm.taobao.org
npm run build:prod
mv dist $BASE/bin/baas-frontend
