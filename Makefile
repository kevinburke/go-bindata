MEGACHECK := $(GOPATH)/bin/megacheck

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
