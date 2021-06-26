#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR="$(readlink -f "$SCRIPT_DIR/..")"
cd "$PROJ_DIR"

PROJECT_NAME="${PROJECT_NAME:-fake-btc-markets}"
VERSION="${VERSION:-$(cat "$PROJ_DIR/VERSION")}"

# Parse VERSION for tags
MAJOR=$(echo "$VERSION" | cut -d '.' -f 1)
MINOR="$MAJOR."$(echo "$VERSION" | cut -d '.' -f 2)

push_tag() {
	image="$1"
	tag="$2"

	echo "Pushing $image:$tag"
	docker tag "$image:$VERSION" "$image:$tag"
	docker push "$image:$tag"
}

push() {
	image="wexel/$PROJECT_NAME-$1"

	push_tag "$image" "$MAJOR"
	push_tag "$image" "$MINOR"
	push_tag "$image" "$VERSION"
	push_tag "$image" latest
}

push api
push db-init
