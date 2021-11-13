#!/bin/bash

echo "Stopping frontend container"

#docker rm -f $(docker ps -aq)
docker-compose stop