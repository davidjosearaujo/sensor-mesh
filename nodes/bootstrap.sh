#!/bin/sh

REPO_NAME=$1        #FIRESTATION_ID:VEHICLE_NAME
EVENT_ID=$2         #Iditifier of the current event
INTF_NAME=$3        #Names of the used network interface

SWARM_KEY=$4        #Private IPFS Swarm key
STORE=$5            #OrbitDB database store address (name if first vehicle)

# Initiate node and clean swarm
ipfs init

#Initiate sensormesh
sensormesh init --nodename=$REPO_NAME --swarmkey=$SWARM_KEY

# Remove bootstrap addresses
ipfs bootstrap rm --all

# Change the routing type to Distributed Hash Table
ipfs config Routing.Type dht

# Starting IPFS daemon service and gives it time to initialize
ipfs daemon --enable-pubsub-experiment &
sleep 5

# Starting the sensormesh daemon
sensormesh daemon --storeaddress=$STORE