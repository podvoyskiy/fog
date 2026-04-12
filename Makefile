VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

build:
	go build -o fog .

release:
	go build $(LDFLAGS) -trimpath -o fog .

clean:
	rm -f fog