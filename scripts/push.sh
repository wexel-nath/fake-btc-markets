#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-fake-btc-markets}"
VERSION="${VERSION-:$(cat "$PROJ_DIR/VERSION")}"

push() {
	image="wexel/$PROJECT_NAME-$1"

	docker tag "$image:$VERSION" "$image:latest"

	echo "Pushing $image:$VERSION"
	docker push "$image:$VERSION"
	docker push "$image:latest"
}

push api
