
#!/bin/bash

export OS_ARCH=$(echo "$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
export PATH=${PWD}/../bin/${OS_ARCH}:$PATH
export FABRIC_CFG_PATH=$PWD/configtx/

export CORE_PEER_TLS_ENABLED=false
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto/crypto-config/peerOrganizations/org1.cathaybc.com/peers/peer0.org1.cathaybc.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto/crypto-config/peerOrganizations/org1.cathaybc.com/users/Admin@org1.cathaybc.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CHAINCODE_NAME='users'

ComplexInputTemplate="{\"key\":\"key\",\"value\":1,\"message\":\"test complex object\"}"

while [[ $# -ge 1 ]] ; do
  key="$1"
  case $key in
  init ) # init
    peer chaincode invoke -o localhost:7050 -C mychannel -n $CHAINCODE_NAME --peerAddresses localhost:7051 --isInit -c '{"function":"InitLedger","Args":[]}'
    shift
    ;;
  1 ) # Query
    peer chaincode query -C mychannel -n $CHAINCODE_NAME -c '{"function":"GetUser","Args":["1"]}'
    peer chaincode query -C mychannel -n $CHAINCODE_NAME -c '{"function":"GetAllUsers","Args":[]}'
    shift
    ;;
  2 ) # Invoke
    peer chaincode invoke -o localhost:7050 -C mychannel -n $CHAINCODE_NAME --peerAddresses localhost:7051 -c '{"function":"CreateUser","Args":["1","Evan","evan@gmail.com"]}'
    peer chaincode invoke -o localhost:7050 -C mychannel -n $CHAINCODE_NAME --peerAddresses localhost:7051 -c '{"function":"CreateUser","Args":["2","Amy","amy@gmail.com"]}'
    peer chaincode invoke -o localhost:7050 -C mychannel -n $CHAINCODE_NAME --peerAddresses localhost:7051 -c '{"function":"DeleteUser","Args":["2"]}'
    shift
    ;;
  * )
    echo
    echo "Unknown flag: $key"
    exit 1
    ;;
  esac
  shift
done


