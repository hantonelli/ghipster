#!/usr/bin/env bash
cd internal/repository/ent/ && go run entgo.io/ent/cmd/entc generate ./schema --target ./entgen