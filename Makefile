# Make 'dis shit!

linter := $(GOPATH)/bin/golint
$(linter):
	go get golang.org/x/lint/golint

SRC_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: lint
lint: linter
	gofmt -d $(SRC_FILES)
	go vet ./...
	$(linter) -set_exit_status ./...

.PHONY: test
test:
	go test -race ./...
