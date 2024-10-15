#!/bin/bash
export $(cat .env | xargs)
tern migrate --migrations ./internal/pgstore/migrations --config ./internal/pgstore/migrations/tern.conf