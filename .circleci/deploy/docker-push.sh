#!/bin/bash

set -e

docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_PASSWORD}

docker tag "statisticopricefinder_rest" "joesweeny/statisticopricefinder_rest:$CIRCLE_SHA1"
docker push "joesweeny/statisticopricefinder_rest:$CIRCLE_SHA1"
