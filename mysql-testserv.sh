#!/bin/bash
# Creates a test MySQL server in Docker with no root password

DOCKER_BIN=$(which docker)

if [ ! -x $DOCKER_BIN ]; then
  echo "Docker binary '$DOCKER_BIN' does not exist, or is not executable!"
  exit 1
fi

$DOCKER_BIN run \
  --detach \
  --name library-api-mysql-testserv \
  --env MYSQL_ALLOW_EMPTY_PASSWORD=true \
  --publish 3306:3306 \
  mysql:latest
