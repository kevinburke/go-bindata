MEGACHECK := $(GOPATH)/bin/megacheck
WRITE_MAILMAP := $(GOPATH)/bin/write_mailmap

all:
	make -C testdata

$(MEGACHECK):
	go get honnef.co/go/tools/cmd/megacheck

lint: | $(MEGACHECK)
	go vet ./...
	$(MEGACHECK) ./...

test:
	go test ./...
	make -C testdata

race-test: lint
	go test -race ./...
	make -C testdata

$(WRITE_MAILMAP):
	go get -u github.com/kevinburke/write_mailmap

force: ;

AUTHORS.txt: force | $(WRITE_MAILMAP)
	$(WRITE_MAILMAP) > AUTHORS.txt

authors: AUTHORS.txt
