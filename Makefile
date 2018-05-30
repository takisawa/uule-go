GO?=go
DEP?=dep

.PHONY: dep depends depends-update test

test:
	$(GO) test ./...

dep:
	$(GO) get -u github.com/golang/dep/cmd/dep

depends:
	$(DEP) ensure -v

depends-update:
	$(DEP) ensure -update
