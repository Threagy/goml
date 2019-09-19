#!/bin/bash

COLUMNS=12
PS3='Please enter your choice: '
options=("Workspace" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Workspace")
            echo "Connect to Workspace"
            docker-compose -f docker-compose.workspace.yml exec workspace bash
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done