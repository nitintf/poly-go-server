#!/bin/sh
set -a && source .env.local && set +a

setup() {

  # Install git hooks
  sh ./scripts/install-git-hooks.sh

  # export DATABASE_URL="postgresql://postgres:postgres@localhost:5432/gograph?sslmode=disable"
  # # 1. Drop DB
  # make db-dump
  # # # 2. Run migrations
  # make db-up
}

setup
