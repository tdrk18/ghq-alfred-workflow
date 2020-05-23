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
	rm -f {github,gitlab,bitbucket,git}.png && \
	convert -background none github.svg github.png && \
	convert -background none gitlab.svg gitlab.png && \
	convert -background none bitbucket.svg bitbucket.png && \
	convert -background none git.svg git.png

.PHONY: build
build:
	mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -o dist/ghq-alfred-workflow

.PHONY: clean
clean:
	cd resources && \
	rm -f {github, gitlab, bitbucket}.png
	cd dist && \
	rm -rf ghq-alfred-workflow ghq-alfred.alfredworkflow resources

.PHONY: distribution
distribution: clean convert build
	mkdir -p dist/resources
	cp resources/*.png dist/resources/
	cd dist && \
	zip -r ghq-alfred.alfredworkflow .

.PHONY: version
version:
	@if [ -z ${VERSION} ]; then \
		echo "usage: make version VERSION='0.1.2'"; \
		exit 1; \
	fi
	plutil -replace version -string ${VERSION} dist/Info.plist
