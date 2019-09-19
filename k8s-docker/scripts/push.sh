#!/bin/bash

REPO=$1
SVC=$2
VER=$3

if [ -z "$REPO" ] || [ -z "$SVC" ] || [ -z "$VER" ];
then
	echo "example:"

	echo "k8s:  make REPO='172.19.148.170:5000' SVC='metalica' VER='v0.5' push"
	echo "vm04:  make REPO='172.20.93.122:5000' SVC='metalica' VER='v0.5' push"

  exit 1
fi

TAG=$REPO/$SVC:$VER

docker tag metalica/$SVC:$VER $TAG
docker push $TAG
docker rmi $TAG