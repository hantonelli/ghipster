#!/usr/bin/env bash

cd backend/ghipster/internal/ent && go run github.com/facebook/ent/cmd/ent generate ./schema --target ./gen