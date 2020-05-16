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

hasCommandByType(){
    if type $1 2>/dev/null; 
    then
        return 1
    else
        return 0
    fi
}
hasCommandByType go
returnVue=$?
echo $returnVue
if [ $returnVue == 0 ] ;
then
    echo "golang runtime is not install"
    rm -rf go*.tar.gz
    wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
    tar -xvzf go1.14.linux-amd64.tar.gz -C /usr/local
    sed -i '/GOROOT/d' /etc/profile
    sed -i '$a\export GOROOT=/usr/local/go' /etc/profile
    sed -i '$a\export PATH=$GOROOT/bin:$PATH' /etc/profile
    source /etc/profile
fi

hasCommandByType node
returnVue=$?
echo $returnVue
if [ $returnVue == 0 ] ;
then
    echo "nodejs runtime is not install"
    rm -rf node*.tar.gz
    wget wget http://nodejs.org/dist/node-latest.tar.gz
    tar -xvzf node-latest.tar.gz -C /usr/local
    sed -i '/NODE_HOME/d' /etc/profile
    sed -i '$a\export NODE_HOME=/usr/local/node' /etc/profile
    sed -i '$a\export PATH=$NODE_HOME/bin:$PATH' /etc/profile
    source /etc/profile
fi

checkEnv(){
    uname -a
    # lsb_release -a
    echo "nodejs runtime:`node -v`"
    echo "nodejs runtime:`go version`"
}
checkEnv

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
