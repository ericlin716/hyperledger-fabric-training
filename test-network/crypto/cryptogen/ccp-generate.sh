#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${MSPID}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ${DIR}/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${MSPID}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ${DIR}/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

DIR=$(dirname "$0")
ORG=1
P0PORT=7051
CAPORT=7054
MSPID=Org1MSP
PEERPEM=crypto/crypto-config/peerOrganizations/org1.cathaybc.com/tlsca/tlsca.org1.cathaybc.com-cert.pem
CAPEM=crypto/crypto-config/peerOrganizations/org1.cathaybc.com/ca/ca.org1.cathaybc.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $MSPID $CAPORT $PEERPEM $CAPEM)" > crypto/crypto-config/peerOrganizations/org1.cathaybc.com/connection-org1.json
echo "$(yaml_ccp $ORG $P0PORT $MSPID $CAPORT $PEERPEM $CAPEM)" > crypto/crypto-config/peerOrganizations/org1.cathaybc.com/connection-org1.yaml
