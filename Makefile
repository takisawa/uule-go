GO?=go
DEP?=dep

.PHONY: dep test

test:
	$(GO) test ./...

dep:
	$(GO) get -u github.com/golang/dep/cmd/dep
