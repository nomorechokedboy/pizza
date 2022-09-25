#!/bin/bash

# !!! Please run this file in root project folder
# This file is a shortcut for `docker-compose up`
# USAGE:
# $ script/compose_start.sh [environment]
# EG:
# $ script/compose_start.sh 
# $ script/compose_start.sh dev

environment=$1

case "$environment" in
    dev)
        docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
        ;;
     
    test)
        docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit
        ;;
    lint)
        docker-compose -f docker-compose.yml -f docker-compose.lint.yml up --abort-on-container-exit
        ;;
    *)
        docker-compose up
esac
