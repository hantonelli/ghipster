#!/usr/bin/env bash
cd internal/repository/ent/ && go run github.com/facebook/ent/cmd/entc generate ./schema --target ./gen