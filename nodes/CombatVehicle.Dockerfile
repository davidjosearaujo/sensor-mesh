FROM ipfs/kubo:latest

ENV VEHICLE_NAME="vfci#"
ENV FIRESTATION_ID="0000"
ENV EVENT_ID="1111"
ENV INTF_NAME="eth0"
ENV SWARM_KEY=""
ENV STORE=""

COPY sensormesh /usr/local/bin/sensormesh
COPY bootstrap.sh /usr/local/bin/bootstrap.sh

RUN chmod -R 777 /usr/local/bin/*

ENTRYPOINT bootstrap.sh ${FIRESTATION_ID}_${VEHICLE_NAME} ${EVENT_ID} ${INTF_NAME} ${SWARM_KEY} ${STORE}