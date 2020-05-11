install-lint:
	@if [ ! -f ./bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.26.0; \
	fi

.PHONY: lint
lint: install-lint
	./bin/golangci-lint run

.PHONY: test
test:
	go test ./... -v -cover
