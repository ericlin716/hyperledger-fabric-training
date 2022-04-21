#設定PATH路徑，為什麼需要設定?請參考 https://iter01.com/558435.html
#如果是Macbook M1架構請改執行 export PATH=${PWD}/../bin/linux-aarch64:$PATH
export PATH=${PWD}/../bin/linux-amd64:$PATH

#移除所有正在執行的container，建議在建置新的區塊鏈網路前執行一次
docker rm -f $(docker ps -qa)

#移除所有volume，建議在建置新的區塊鏈網路前執行一次
docker volume prune -f

#建議每次重新架鏈前都刪掉之前建立的憑證
rm -rf ./crypto-config/

#建議每次重新架鏈前都刪掉之前鏈相關的channel文件
rm -rf ./channel-artifacts/

#################################### 以下為 Fabric 相關的指令 ####################################

#產生orderer憑證
cryptogen generate --config=./config/crypto-config-orderer.yaml --output=./crypto-config

#產生peer憑證
cryptogen generate --config=./config/crypto-config-org1.yaml --output=./crypto-config

#產生創世區塊
configtxgen -configPath ./config -profile OrderersGenesis -channelID system-channel -outputBlock ./system-genesis-block/genesis.block

#產生create application channel文件
configtxgen -configPath ./config -profile OrgsChannel -outputCreateChannelTx ./channel-artifacts/cathay.tx -channelID cathay

#產生update anchor peer文件
configtxgen -configPath ./config -profile OrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/org1Anchors.tx -channelID cathay -asOrg Org1MSP

#進入docker container
docker exec -it peer1.org1.cathaybc.com sh


#################################### 以下為 container 內的指令 ####################################

#設定使用admin憑證
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/crypto-config/peerOrganizations/org1.cathaybc.com/users/Admin@org1.cathaybc.com/msp/

#建立channel
peer channel create -o orderer1.cathaybc.com:7050 -c cathay -f ./channel-artifacts/cathay.tx --outputBlock ./channel-artifacts/cathay.block --tls --cafile ./crypto-config/ordererOrganizations/cathaybc.com/orderers/orderer1.cathaybc.com/tls/ca.crt

#加入channel
peer channel join -b ./channel-artifacts/cathay.block

#更新anchor peer
peer channel update -o orderer1.cathaybc.com:7050 -c cathay -f ./channel-artifacts/org1Anchors.tx --tls --cafile ./crypto-config/ordererOrganizations/cathaybc.com/orderers/orderer1.cathaybc.com/tls/ca.crt

#打包
peer lifecycle chaincode package test-chaincode.tar.gz --path /chaincode/test-chaincode --lang golang --label test-chaincode_1

#安裝
peer lifecycle chaincode install test-chaincode.tar.gz

#查看安裝
peer lifecycle chaincode queryinstalled

#批准
peer lifecycle chaincode approveformyorg -o orderer1.cathaybc.com:7050 --channelID cathay --name test-chaincode --version 1 --sequence 1 --init-required --package-id [貼上 Package ID] --tls --cafile ./crypto-config/ordererOrganizations/cathaybc.com/orderers/orderer1.cathaybc.com/tls/ca.crt

#查看批准狀態
peer lifecycle chaincode checkcommitreadiness --channelID cathay --name test-chaincode --version 1 --sequence 1 --output json --init-required

#提交
#若需要兩節點背書，須加上 --peerAddresses peer1.org1.cathaybc.com:7051 --tlsRootCertFiles ./crypto-config/peerOrganizations/org1.cathaybc.com/peers/peer1.org1.cathaybc.com/tls/ca.crt --peerAddresses peer1.org2.cathaybc.com:7051 --tlsRootCertFiles ./crypto-config/peerOrganizations/org2.cathaybc.com/peers/peer1.org2.cathaybc.com/tls/ca.crt
peer lifecycle chaincode commit -o orderer1.cathaybc.com:7050 --channelID cathay --name test-chaincode --version 1 --sequence 1 --init-required --waitForEvent --tls --cafile ./crypto-config/ordererOrganizations/cathaybc.com/orderers/orderer1.cathaybc.com/tls/ca.crt

#查看提交
peer lifecycle chaincode querycommitted --channelID cathay --name test-chaincode

#初始化
#若需要兩節點背書，須加上 --peerAddresses peer1.org1.cathaybc.com:7051 --tlsRootCertFiles ./crypto-config/peerOrganizations/org1.cathaybc.com/peers/peer1.org1.cathaybc.com/tls/ca.crt --peerAddresses peer1.org2.cathaybc.com:7051 --tlsRootCertFiles ./crypto-config/peerOrganizations/org2.cathaybc.com/peers/peer1.org2.cathaybc.com/tls/ca.crt
peer chaincode invoke -o orderer1.cathaybc.com:7050 -C cathay -n test-chaincode --isInit -c '{"function":"Init","Args":[]}' --waitForEvent --tls --cafile ./crypto-config/ordererOrganizations/cathaybc.com/orderers/orderer1.cathaybc.com/tls/ca.crt

#Query example
peer chaincode query -C cathay -n test-chaincode -c '{"function":"QueryFunction1","Args":[]}'

#Invoke example
#若需要兩節點背書，須加上 --peerAddresses peer1.org1.cathaybc.com:7051 --tlsRootCertFiles ./crypto-config/peerOrganizations/org1.cathaybc.com/peers/peer1.org1.cathaybc.com/tls/ca.crt --peerAddresses peer1.org2.cathaybc.com:7051 --tlsRootCertFiles ./crypto-config/peerOrganizations/org2.cathaybc.com/peers/peer1.org2.cathaybc.com/tls/ca.crt
peer chaincode invoke -o orderer1.cathaybc.com:7050 -C cathay -n test-chaincode -c '{"function":"InvokeFunction1","Args":[]}' --waitForEvent --tls --cafile ./crypto-config/ordererOrganizations/cathaybc.com/orderers/orderer1.cathaybc.com/tls/ca.crt

#################################### 以下為其他指令 ####################################

#查看憑證
openssl x509 -noout -text -in [CERT PATH]