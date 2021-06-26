#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-fake-btc-markets}"
VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

build() {
	image="wexel/$PROJECT_NAME-$1"
	dockerfile="./docker/Dockerfile.$1"

	docker build \
		-t "$image:$VERSION" \
		-f "$dockerfile" \
		.
}

build api
build db-init
