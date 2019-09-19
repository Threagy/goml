#!/bin/bash

COLUMNS=12
PS3='Please enter your choice: '
options=(
          "Workspace [Daemon]"
          "Gomlet [Production]"
          "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Workspace [Daemon]")
            echo "Workspace up..."
            docker-compose -f docker-compose.workspace.yml up -d
            break
            ;;
        "Gomlet [Production]")
            echo "Gomlet [Production] up..."
            docker-compose -f docker-compose.workspace.prod.yml up -d
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done