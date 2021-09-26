#!/bin/bash

echo "Stopping nlp engine container"

#docker rm -f $(docker ps -aq)
docker-compose down