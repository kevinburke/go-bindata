DIFFER := $(GOPATH)/bin/differ
MEGACHECK := $(GOPATH)/bin/megacheck
WRITE_MAILMAP := $(GOPATH)/bin/write_mailmap

all:
	$(MAKE) -C testdata

$(DIFFER):
	go get -u github.com/kevinburke/differ

diff-testdata: | $(DIFFER)
	$(DIFFER) $(MAKE) -C testdata
	$(DIFFER) go fmt ./testdata/out/...

$(MEGACHECK):
	go get honnef.co/go/tools/cmd/megacheck

lint: | $(MEGACHECK)
	go vet ./...
	$(MEGACHECK) ./...

go-test:
	go test ./...

go-race-test:
	go test -race ./...

test: go-test
	$(MAKE) -C testdata

race-test: lint go-race-test
	$(MAKE) -C testdata

$(WRITE_MAILMAP):
	go get -u github.com/kevinburke/write_mailmap

force: ;

AUTHORS.txt: force | $(WRITE_MAILMAP)
	$(WRITE_MAILMAP) > AUTHORS.txt

authors: AUTHORS.txt

ci: go-race-test diff-testdata
