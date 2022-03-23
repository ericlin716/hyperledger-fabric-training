#!/bin/bash

export OS_ARCH=$(echo "$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
export PATH=${PWD}/../bin/${OS_ARCH}:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx

VERBOSE=false
MAX_RETRY=5
# default for delay between commands
CLI_DELAY=3
# channel name defaults to "mychannel"
CHANNEL_NAME="mychannel"
# use golang as the default language for chaincode
CC_SRC_LANGUAGE=golang
# Chaincode name
CHAINCODE_NAME='test-chaincode'

while [ -z "$VERSION" ]; do
    read -p "Input chaincode version: " VERSION
done


scripts/deployCC.sh $CHANNEL_NAME $CC_SRC_LANGUAGE $VERSION $CHAINCODE_NAME $CLI_DELAY $MAX_RETRY $VERBOSE
if [ $? -ne 0 ]; then
    echo "ERROR !!! Deploying chaincode failed"
fi