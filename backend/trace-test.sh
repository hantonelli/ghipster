#!/bin/bash

curl \
-X POST \
-H "Content-Type: application/json" \
-H "x-request-id: dasdads" \
--data '{"operationName":null,"variables":{},"query":"mutation {\n  createProduct(input: {name: \"product2\"}) {\n    id\n    name\n  }\n}\n"}' \
http://localhost:8080/query