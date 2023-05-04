TEST?=$$(go list ./... |grep -v 'vendor')
TEST_PARALLELISM?=4
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=cors

default: build

build:
	go install $(FLAGS)
	
test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=$(TEST_PARALLELISM)

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 360m -parallel $(TEST_PARALLELISM)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: build test testacc vet fmt test-compile
