#!/bin/bash

VERSION=v1
OUTPUT_PATH=./proto/$VERSION
PROTOC=/usr/share/protobuf/bin/protoc

# Create Directory
if [ ! -d "$OUTPUT_PATH" ]; then
  mkdir -p $OUTPUT_PATH
fi

####################################################
# common

TARGET=common

# generate proto files
declare -a protofiles=(
  "./proto/$VERSION/$TARGET/docker.proto"
  "./proto/$VERSION/$TARGET/node.proto")

for protofile in "${protofiles[@]}"""; do
  filename=$(basename -- "$protofile")
  mkdir -p ./proto/$VERSION/$TARGET/ && cp ../proto/$VERSION/$TARGET/$filename ./proto/$VERSION/$TARGET/$filename
  echo ./proto/$VERSION/$TARGET/$TARGET/$filename "copied."
done

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=plugins=grpc,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include -I. \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --swagger_out=logtostderr=true:. \
  ${protofiles[@]}

####################################################
# gomlapi

TARGET=gomlapi

# generate proto files
declare -a protofiles=(
  "./proto/$VERSION/$TARGET/gomlapi.proto"
  "./proto/$VERSION/$TARGET/gomlapi-common.proto")

for protofile in "${protofiles[@]}"""; do
  filename=$(basename -- "$protofile")
  mkdir -p ./proto/$VERSION/$TARGET/ && cp ../proto/$VERSION/$TARGET/$filename ./proto/$VERSION/$TARGET/$filename
  echo ./proto/$VERSION/$TARGET/$TARGET/$filename "copied."
done

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=plugins=grpc,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include -I. \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --swagger_out=logtostderr=true:. \
  ${protofiles[@]}

####################################################
# gomlet

TARGET=gomlet

# generate proto files
declare -a protofiles=(
  "./proto/$VERSION/$TARGET/gomlet.proto"
  "./proto/$VERSION/$TARGET/gomlet-common.proto")

for protofile in "${protofiles[@]}"""; do
  filename=$(basename -- "$protofile")
  mkdir -p ./proto/$VERSION/$TARGET/ && cp ../proto/$VERSION/$TARGET/$filename ./proto/$VERSION/$TARGET/$filename
  echo ./proto/$VERSION/$TARGET/$TARGET/$filename "copied."
done

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=plugins=grpc,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include -I. \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --swagger_out=logtostderr=true:. \
  ${protofiles[@]}

####################################################
# auth

TARGET=auth

# generate proto files
declare -a protofiles=(
  "./proto/$VERSION/$TARGET/auth.proto"
  "./proto/$VERSION/$TARGET/auth-common.proto")

for protofile in "${protofiles[@]}"""; do
  filename=$(basename -- "$protofile")
  mkdir -p ./proto/$VERSION/$TARGET/ && cp ../proto/$VERSION/$TARGET/$filename ./proto/$VERSION/$TARGET/$filename
  echo ./proto/$VERSION/$TARGET/$TARGET/$filename "copied."
done

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=plugins=grpc,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include -I. \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --swagger_out=logtostderr=true:. \
  ${protofiles[@]}

####################################################
# user

TARGET=user

# generate proto files
declare -a protofiles=(
  "./proto/$VERSION/$TARGET/user.proto"
  "./proto/$VERSION/$TARGET/user-common.proto")

for protofile in "${protofiles[@]}"""; do
  filename=$(basename -- "$protofile")
  mkdir -p ./proto/$VERSION/$TARGET/ && cp ../proto/$VERSION/$TARGET/$filename ./proto/$VERSION/$TARGET/$filename
  echo ./proto/$VERSION/$TARGET/$TARGET/$filename "copied."
done

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=plugins=grpc,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
  ${protofiles[@]}

$PROTOC -I/usr/local/include -I. \
  -I. \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
  --swagger_out=logtostderr=true:. \
  ${protofiles[@]}

echo 'proto-gen success.'
