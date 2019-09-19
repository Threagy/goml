#!/bin/bash

COLUMNS=12
PS3='Please enter your choice: '
options=(
          "Workspace"
          "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Workspace")
            echo "Workspace down..."
            docker-compose -f docker-compose.workspace.yml down
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done