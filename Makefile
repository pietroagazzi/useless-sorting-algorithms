.PHONY: build test clean

GO=go
BINARY=bin/useless-sorting-algorithms
BINARY_PATH=cmd/cli/main.go

build:
	$(GO) build -o $(BINARY) $(BINARY_PATH)

test:
	$(GO) test -v ./pkg/...

run : build
	@echo "Running ${BINARY} with args: ${RUN_ARGS}"

	./${BINARY} ${ARGS}

clean:
	rm -f $(BINARY)	