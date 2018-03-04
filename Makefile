# Make 'dis shit!

linter := $(GOPATH)/bin/golint
$(linter):
	go get golang.org/x/lint/golint

SRC_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

guru := $(GOPATH)/bin/guru
$(guru):
	go get golang.org/x/tools/cmd/guru


# tools are all external commands used by gomate
tools: $(guru) $(linter)

# assets folder contents are bundles with go executable
cmd/gomate/assets.go: assets/*
	go get github.com/pokstad/go-bindata/...
	go-bindata -o cmd/gomate/assets.go assets

.PHONY: lint
lint: $(linter)
	gofmt -d $(SRC_FILES)
	go vet ./...
	$(linter) -set_exit_status ./...

.PHONY: test
test: tools
	go test -race ./...
