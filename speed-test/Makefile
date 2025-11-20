os-archs ?= \
	android:arm64 \
	darwin:amd64 \
	darwin:arm64 \
	linux:amd64 \
	linux:arm64 \
	windows:amd64 \
	windows:arm64

.PHONY: build
build:
	@module="$$(go list -m)"; \
	echo "$$module"; \
	for os_arch in $(os-archs); do \
		os=$$(echo $$os_arch | cut -d: -f1); \
		arch=$$(echo $$os_arch | cut -d: -f2); \
		out="$$module-$$os-$$arch"; \
		echo "Building $$out"; \
		if [ "$$os" = "windows" ]; then \
			out="$$out.exe"; \
		fi; \
		GOOS=$$os \
		GOARCH=$$arch \
		go build -o ./dist/server-$$out ./server/main.go; \
		go build -o ./dist/client-$$out ./client/main.go; \
	done