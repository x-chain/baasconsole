# baasmanager

![](https://img.shields.io/badge/build-passing-brightgreen.svg)
![](https://img.shields.io/badge/author-jonluo-yellow.svg)
![](https://img.shields.io/badge/kubernetes-v1.14.1-blue.svg)
![](https://img.shields.io/badge/go-v1.12.5-blue.svg)
![](https://img.shields.io/badge/docker-v18.06.3–ce-blue.svg)
![](https://img.shields.io/badge/hyperledger fabric-v1.4.1-blue.svg)

### 基于 K8S 平台的区块链即服务（Blockchain as a Service）

### 整体功能

#### 动态创建 fabric

- [x] solo
- [x] kafka
- [x] etcdraft

#### 区块链监控

- [x] 区块链首页统计分析
- [x] 区块链浏览器

#### 区块链资源

- [x] 动态扩容
- [x] 释放

### 主要目录结构

- baas-kubecluster  
  k8s 集群，基于 flannel 网络，安装 dashboard 插件，还有其余插件等 (一个简单的 k8s 集群)
- baas-nfsshared  
  其会生成 baas-artifacts，baas-fabric-data，baas-k8s-config 目录
  - baas-artifacts 为存放生成的证书文件
  - baas-fabric-data 为 fabric 网络映射出来的数据
  - baas-k8s-config 为生成的 k8s yaml 定义文件
- baas-template  
  fabric k8s 的模板文件，用于生成 baas-nfsshared/baas-k8s-config 下的文件
- baas-fabricengine  
  用于生成 baas-nfsshared 的文件即目录结构和执行 fabric 操作
- baas-kubeengine  
  kubeconfig/config 文件是 k8s master 的\$HOME/.kube/config 文件，用于 k8s client 链接 k8s 集群，将 baas-nfsshared/baas-k8s-config 下的文件在 k8s 集群创建启动
- baas-gateway  
  统一 api 网关管理，调用入口
- baas-frontend  
  baas admin 前端

### 架构图

![](baas-others/images/baas.png)

### 数据流图

![](baas-others/images/flow.png)

### 页面

#### 首页信息

![](baas-others/images/das.png)

#### 用户管理

![](baas-others/images/user.png)

#### 角色管理

![](baas-others/images/role.png)

#### 区块链管理

![](baas-others/images/chain.png)

#### 通道管理

![](baas-others/images/channel.png)

#### 合约管理和区块链浏览器

![](baas-others/images/chaincode.png)

#### 智能合约调用，查询

![](baas-others/images/cc1.png)

#### 智能合约调用，调用

![](baas-others/images/cc2.png)

### 部署样例

- [简单部署样例](sample.md)
