#!/bin/bash

# This script gets the AWS Instance Metadata Simulator.

yum install git golang
mkdir ~/go
export GOPATH=~/go
go get -u github.com/mtnfog/aws-metadata-emulator
cd ~/go/bin
wget https://raw.githubusercontent.com/mtnfog/aws-metadata-simulator/master/metadata.toml
sudo iptables -t nat -A OUTPUT -p tcp -d 169.254.169.254 --dport 80 -j DNAT --to-destination 127.0.0.1:8080
echo "Edit the metadata.toml to have the values you need and then run ./aws-metadata-simulator"
