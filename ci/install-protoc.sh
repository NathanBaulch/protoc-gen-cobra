#!/bin/bash

set -ex

curl -sSL "https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.12.4-linux-x86_64.zip" -o /tmp/protoc.zip
unzip /tmp/protoc.zip -d /tmp/protoc
sudo mv /tmp/protoc/bin/protoc /usr/local/bin/protoc
