#!/bin/bash

# https://github.com/swagger-api/swagger-ui/issues/3675

LOCAL_PORT=87
LOCAL_SWAGGER_FILE=/c/Users/horacio.antonelli/projects/ghipster/backend/third_party/swagger.json

docker run \
    -p ${LOCAL_PORT}:8080 \
    -e SWAGGER_JSON=/usr/share/nginx/html/swagger.json \
    -v ${LOCAL_SWAGGER_FILE}:/usr/share/nginx/html/swagger.json \
    swaggerapi/swagger-ui 