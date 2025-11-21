os-archs ?= \
	android:arm64 \
	darwin:amd64 \
	darwin:arm64 \
	linux:amd64 \
	linux:arm64 \
	windows:amd64 \
	windows:arm64
module = $(shell go list -m)

.PHONY: build
build:
	for os_arch in $(os-archs); do \
		os=$$(echo $$os_arch | cut -d: -f1); \
		arch=$$(echo $$os_arch | cut -d: -f2); \
		out="$(module)-$$os-$$arch"; \
		ext="" \
		[ "$$os" = "windows" ] && ext=".exe"; \
		echo "Building $$out"; \
		CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch \
			go build -trimpath -ldflags "-s -w" -o ./dist/$$out-server$$ext ./server/main.go; \
		CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch \
			go build -trimpath -ldflags "-s -w" -o ./dist/$$out-client$$ext ./client/main.go; \
	done
