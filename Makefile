.PHONY: docker run

BINARY=main
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
changelog_args=-o CHANGELOG.md -tag-filter-pattern '^v'

check-modd-exists:
	@modd --version > /dev/null

run: check-modd-exists
	@modd -f ./.modd/server.modd.conf

run-worker: check-modd-exists
	@modd -f ./.modd/worker.modd.conf

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

check-cognitive-complexity:
	find . -type f -name '*.go' -not -name "mock*.go" -not -name "generated.go" -not -path "./api/v1/*"  -not -path "./scripts/*"\
      -exec gocognit -over 20 {} +

clean:
	rm -v internal/entity/mock/mock_*.go

check-gotest:
ifeq (, $(shell which richgo))
	$(warning "richgo is not installed, falling back to plain go test")
	$(eval TEST_BIN=go test)
else
	$(eval TEST_BIN=richgo test)
endif

ifdef test_run
	$(eval TEST_ARGS := -run $(test_run))
endif
	$(eval test_command=$(TEST_BIN) ./... $(TEST_ARGS) -v --cover)

changelog:
ifdef version
	@echo "Updating version in config.yml to $(version)"
	sed -i 's/version: .*/version: "$(subst v,,$(version))"/' config.yml
	@echo "Updating version in config.prod.yml to $(version)"
	sed -i 's/version: .*/version: "$(subst v,,$(version))"/' config.prod.yml
endif
ifdef version
	$(eval changelog_args=--next-tag $(version) $(changelog_args))
endif
	git-chglog $(changelog_args)
