MEGACHECK := $(GOPATH)/bin/megacheck

BAZEL_VERSION := 0.7.0
BAZEL_DEB := bazel_$(BAZEL_VERSION)_amd64.deb

all:
	make -C testdata

$(MEGACHECK):
	go get honnef.co/go/tools/cmd/megacheck

lint: | $(MEGACHECK)
	go vet ./...
	$(MEGACHECK) ./...

install-travis:
	wget "https://storage.googleapis.com/bazel-apt/pool/jdk1.8/b/bazel/$(BAZEL_DEB)"
	sudo dpkg --force-all -i $(BAZEL_DEB)
	sudo apt-get install moreutils -y

test:
	bazel test \
		--experimental_repository_cache="$$HOME/.bzrepos" \
		--test_output=errors //...
	$(MAKE) -C testdata

ci: lint
	bazel --batch --host_jvm_args=-Dbazel.DigestFunction=SHA1 test \
		--experimental_repository_cache="$$HOME/.bzrepos" \
		--spawn_strategy=remote \
		--test_output=errors \
		--strategy=Javac=remote \
		--profile=profile.out \
		--noshow_progress \
		--noshow_loading_progress \
		--features=race //...
	$(MAKE) -C testdata

race-test: lint
	bazel test \
		--experimental_repository_cache="$$HOME/.bzrepos" \
		--test_output=errors --features=race //...
	$(MAKE) -C testdata
