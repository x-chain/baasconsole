#!/bin/bash
# 当前目录
export BASE=$(cd `dirname "${BASH_SOURCE[0]}"`;pwd) 
echo $BASE
buildOutput=$BASE/build
execDate=`date +%Y%m%d`

if [ ! -d "$buildOutput" ];then
mkdir $buildOutput
else
rm -rf $buildOutput/*
fi

echo "编译baas-gateway"
cd $BASE/baas-gateway
go build .
mv baas-gateway $buildOutput
cp *.yaml $buildOutput/

echo "编译baas-fabricengine"
cd $BASE/baas-fabricengine
go build .
mv baas-fabricengine $buildOutput
cp *.yaml $buildOutput

echo "编译baas-kubeengine"
cd $BASE/baas-kubeengine
go build .
mv baas-kubeengine $buildOutput
cp *.yaml $buildOutput

echo "编译baas-frontend"
cd $BASE/baas-frontend
rm -rf node_modules && npm install --registry=https://registry.npm.taobao.org
npm run build:prod
mv dist $buildOutput/baas-frontend
cp -r etc $buildOutput
