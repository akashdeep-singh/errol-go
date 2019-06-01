PKGS := $(shell go list ./pkg/...)

.PHONY: build
build: test
	go build

.PHONY: test
test:
	go test $(PKGS)
