GOARCH=amd64
PLATFORMS=darwin linux windows
PKGS=$(shell go list ./...)
BIN_DIR=$(GOPATH)/bin
BINARY=gh-stars
VERSION=$(shell grep "const Version" gh-stars.go | sed 's/[^0-9\.]//g')

build:
	go build -o $(BIN_DIR)/$(BINARY) .
	chmod +x $(BIN_DIR)/$(BINARY)

clean:
	rm -rf dist

os=$(word 1,$@)
.PHONY: ($PLATFORMS)
$(PLATFORMS):
	@mkdir -p dist > /dev/null
	GOOS=$(os) GOARCH=amd64 go build -o dist/$(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: clean windows linux darwin
