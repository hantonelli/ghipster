#!/usr/bin/env bash
cd internal/service/ent/ && go run entgo.io/ent/cmd/entc generate ./schema --target ./gen