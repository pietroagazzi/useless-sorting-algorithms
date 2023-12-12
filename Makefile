.PHONY: build test clean

GO=go
BINARY=bin/uselessSortingAlgorithms
BINARY_PATH=cmd/uselessSortingAlgorithms/main.go

build:
	$(GO) build -o $(BINARY) $(BINARY_PATH)

test:
	$(GO) test ./...

run : build
	@echo "Running ${BINARY} with args: ${RUN_ARGS}"

	./${BINARY} ${ARGS}

clean:
	rm -f $(BINARY)	