#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-fake-btc-markets}"
export VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

deploy() {
	docker stack deploy \
		--compose-file 'docker/docker-stack.yml' \
		"$PROJECT_NAME"
}

deploy
