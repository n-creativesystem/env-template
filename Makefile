NAME	:= kubectl-vartpl
SOURCES := $(shell find . -name '*.go')
LDFLAGS	:= -ldflags="-s -w -extldflags \"-static\""
GOOS	?= linux
GOARCH	?= amd64

build: kubectl-vartpl

$(NAME): $(SOURCES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(GOOS)-$(GOARCH)/$(NAME)
