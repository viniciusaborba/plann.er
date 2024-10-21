package main

//go:generate goapi-gen --package=planner --out ./internal/api/spec/planner.gen.spec.go ./internal/api/spec/planner.spec.json
//go:generate ./migrate.sh
//go:generate sqlc generate -f ./internal/pgstore/sqlc.yaml

