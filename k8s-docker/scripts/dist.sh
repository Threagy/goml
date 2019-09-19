#!/bin/bash

if [ ! -d "./dist" ]; then
  mkdir ./dist
fi

# gomlet
docker-compose -f docker-compose.workspace.yml exec workspace bash -c 'cd goml/gomlet && make dist'
if [ $? -eq 0 ]; then
    echo "gomlet success"
else
    echo "fail. is running a container??"
    exit 1
fi

# gomlweb
docker-compose -f docker-compose.workspace.yml exec workspace bash -c 'cd goml/gomlweb && make dist'
if [ $? -eq 0 ]; then
    echo "gomlweb success"
else
    echo "fail. is running a container??"
    exit 1
fi

echo "Completed."