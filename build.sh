#!/bin/bash

#   Usage:
##  $ ./build.sh <docker image name> [OPTIONS]
##
##  OPTIONS:
##      --build         - Builds a new image with the specified name
##      --up            - Starts containers using the image with the specified name

TARGET_IMAGE_NAME=$1

if [[ $2 == "--build" ]]; then

    # Build golang configuration file
    cd sensormesh/
    OUT=$(go build -o ../nodes/sensormesh main.go)
    if [[ $OUT != "" ]]; then
        echo -e "[!] sensormesh compiling error !"
    fi
    echo -e "[🗸] SensorMesh binary compiled!"
    cd ..

    # Guarantees that every shell script is in Unix format. (Usefull if you develop in Windows)
    dos2unix -q nodes/*.sh

    # Build new image
    echo -e "[*] Building docker image..."
    SHA=$(docker build -q --no-cache -t $TARGET_IMAGE_NAME -f nodes/CombatVehicle.Dockerfile ./nodes) && echo -e "[🗸] Docker image built. $SHA"
fi

if [[ $2 == "--up" || $3 == "--up" ]]; then
    #  These keys will be created and hand out by a third party that creates the events !!

    ## Generate Swarm key wich must be shared between all nodes
    SWARM_KEY=$(head -c 32 /dev/urandom | od -t x1 -A none - | tr -d '\n '; echo '')
    echo -e "[🗸] Swarm key generated: $SWARM_KEY"

    # First node creates the OrbitDB store with default name
    echo -e "[!] First vehicle on scene..."
    TARGET_IMAGE_NAME=$TARGET_IMAGE_NAME SWARM_KEY=$SWARM_KEY STORE=$STORE docker-compose -f docker-compose-initial.yml --log-level ERROR up -d --remove-orphans
    sleep 10

    # Search for the OrbitDB address inside first node
    for LID in $(docker container ls -q); do
        STORE=$(docker exec $LID sensormesh config orbitdb.storeaddress)
        if [[ $STORE == "" ]]; then
            echo -e "[!] Node failed to run sensormesh!"
            return
        fi
        echo -e "[🗸] Store address created by first vehicle: $STORE"
        break
    done
    sleep 2

    # Spin up secondary nodes with created database address to connect to 
    echo -e "[!] Other vehicle arriving on scene..."
    TARGET_IMAGE_NAME=$TARGET_IMAGE_NAME SWARM_KEY=$SWARM_KEY STORE=$STORE docker-compose -f docker-compose-secondary.yml --log-level ERROR up -d
fi