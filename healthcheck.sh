#!/bin/bash
CONTAINER_NAME=$1
PORT=$2
# healthcheck script checks if the container is running
# on the valid port by checking the port mapping
if [ -z CONTAINER_NAME ]; then
  echo "No container name specified";
  exit 1;
fi
docker-compose ps --filter status=running $CONTAINER_NAME | grep $2/tcp
if [ $? -ne 0 ]; then
  echo "Container ${CONTAINER_NAME} not in healthy state"
  exit 1;
fi
echo "Health check passed"
exit 0
