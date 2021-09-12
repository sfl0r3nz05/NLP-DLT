#!/bin/bash

echo "Stopping nlp engine containers"

#docker rm -f $(docker ps -aq)
docker-compose down