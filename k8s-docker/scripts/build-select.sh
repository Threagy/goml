#!/bin/bash

COLUMNS=12
PS3='Please enter your choice: '
options=(
          "Workspace"
          "Gomlet [Production]"
          "Gomlweb [Production]"
          "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Workspace")
            echo "Workspace build..."
            docker-compose -f docker-compose.workspace.yml build
            break
            ;;
        "Gomlet [Production]")
            echo "Gomlet [Production] build..."
            docker-compose -f docker-compose.gomlet.prod.yml build
            break
            ;;
        "Gomlweb [Production]")
            echo "Gomlet [Production] build..."
            docker-compose -f docker-compose.gomlweb.prod.yml build
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done