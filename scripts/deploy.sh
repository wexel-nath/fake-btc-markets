#!/bin/bash
set -e

PROJECT_NAME="${PROJECT_NAME:-fake-btc-markets}"

deploy() {
	docker stack deploy \
		--compose-file 'docker/docker-stack.yml' \
		"$PROJECT_NAME"
}

deploy
