#!/usr/bin/env bash
docker build -t localhost:32000/graphql:registry -f Dockerfile.graphql.dev .
# docker push localhost:32000/graphql:registry
# docker build -t graphql -f Dockerfile.graphql.dev .
# docker tag 1fe3d8f47868 localhost:3000/graphql:registry