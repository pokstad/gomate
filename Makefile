# Make 'dis shit!

linter := $(GOPATH)/bin/gometalinter.v2
$(linter):
	go get gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

SRC_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

guru := $(GOPATH)/bin/guru
$(guru):
	go get golang.org/x/tools/cmd/guru

gogetdoc := $(GOPATH)/bin/gogetdoc
$(gogetdoc):
	go get github.com/zmb3/gogetdoc

gocode := $(GOPATH)/bin/gocode
$(gocode):
	go get github.com/stamblerre/gocode

# tools are all external commands used by gomate
tools: $(guru) $(linter) $(gogetdoc) $(gocode)

# assets folder contents are bundles with go executable
cmd/gomate/assets.go: assets/*
	go get github.com/pokstad/go-bindata/...
	go-bindata -o cmd/gomate/assets.go assets

$(GOPATH)/bin/gomate: cmd/gomate/assets.go
	go install ./cmd/gomate

install: $(GOPATH)/bin/gomate tools

update: cmd/gomate/assets.go README.md

.PHONY: lint
lint: $(linter)
	$(linter) --config .gometalinter.json ./...

.PHONY: test
test: tools
	go test -race ./...
