#!/bin/bash

echo "Stopping backend container"

#docker rm -f $(docker ps -aq)
docker-compose down