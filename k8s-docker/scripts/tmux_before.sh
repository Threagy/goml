#!/bin/bash

containerIDs=($(docker-compose -f docker-compose.workspace.yml ps -q))

# 컨테이너가 없다면 실행한다.
if [ "${containerIDs[0]}" == "" ] ; then
    echo create workspace container...
    docker-compose -f docker-compose.workspace.yml up -d
    exit 0
fi

# 컨테이너가 멈춰있다면 실행한다.
status=($(docker inspect --format='{{.State.Status}}' ${containerIDs[0]}))
if [ "$status" == "exited" ] ; then
    echo run workspace container...
    docker-compose -f docker-compose.workspace.yml up -d
    exit 0
fi

echo workspace is running...