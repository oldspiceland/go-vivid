SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

# Name for the binary output
BINARY=go-vivid

# Values passed to compiler at compile-time
VERSION=0.0
BUILD_TIME=`date +%FT%T%z`

# Setup the ldflags for compilation via go build, interpolate the variables
LDFLAGS=-ldflags "-X github.com/oldspiceland/go-vivid/main.Version=${VERSION} -X github.com/oldspiceland/go-vivid/main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
	go build ${LDFLAGS} -o ${OUTPUT} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
