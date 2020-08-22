
1, depend docker:
```dotnetcli
hyperledger/fabric-ca:1.4.1
hyperledger/fabric-tools:1.4.1
busybox:latest
hyperledger/fabric-kafka:0.4.15
hyperledger/fabric-orderer:1.4.1
hyperledger/fabric-couchdb:0.4.15
hyperledger/fabric-peer:1.4.1
hyperledger/fabric-zookeeper:0.4.15
```

2, pull images:
cd baas-template
grep image *.yaml | awk -F': ' '{print $3}' | xargs -I@ docker pull @

3,
