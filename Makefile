TESTDIR := testdir

$(TESTDIR):
	mkdir -p $(TESTDIR)/{data,cache}

install-lint:
	@if [ ! -f ./bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.26.0; \
	fi

.PHONY: lint
lint: install-lint
	./bin/golangci-lint run

.PHONY: test
test: $(TESTDIR)
	env alfred_workflow_bundleid=testid \
		alfred_workflow_cache=$(TESTDIR)/cache \
		alfred_workflow_data=$(TESTDIR)/data \
		go test ./... -v -cover

convert:
	cd resources && \
	rm -f {github, gitlab, bitbucket}.png && \
	convert -background none github.svg github.png && \
	convert -background none gitlab.svg gitlab.png && \
	convert -background none bitbucket.svg bitbucket.png && \
	convert -background none git.svg git.png
