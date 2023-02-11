#!/usr/bin/env bash

# Install pre commit & required dependencies.
set -eu
set -o pipefail 

echo "--Installing git-hooks & dependencies---"
if ! hash golangci-lint > /dev/null; then
    echo "golangci-lint not found. installing ..."
    brew install golangci-lint
fi

touch  .git/hooks/pre-commit
echo "golangci-lint run" >> .git/hooks/pre-commit
chmod 755 .git/hooks/pre-commit
