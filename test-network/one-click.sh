#!/bin/bash

./network.sh gencert
./network.sh down
./network.sh up
./network.sh createChannel
./network.sh deployCC
