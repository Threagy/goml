#!/bin/bash

rm -rf /dist/gomlweb
mkdir /dist/gomlweb

rm -rf ./dist

npm run build
cd ./dist
tar -czf /dist/gomlweb/gomlweb.tar.gz .
