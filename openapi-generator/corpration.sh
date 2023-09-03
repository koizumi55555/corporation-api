#!/bin/bash
mkdir -p ${PWD}/internal/controller/http/model
docker run --rm \
    -v ${PWD}/openapi/corporation-openapi:/local \
    -v ${PWD}/internal/controller/http:/project \
    openapitools/openapi-generator-cli:v5.2.0 generate \
    -i /local/openapi.yml \
    -c /project/.openapi-gen.json \
    -g go-server \
    -o /project/model \
    --ignore-file-override=/project/.openapi-generator-ignore \
    --global-property model
