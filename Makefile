HARMONY_GIT_OPTS?=-b v4.3.1
BLS_GIT_OPTS?=-b v0.0.6
MCL_GIT_OPTS?=
GIT_OPTS?=--depth 1
DEBUG_RETRY?=3
ABIGEN?=abigen

gopath=$(shell go env GOPATH)
harmony_root=${gopath}/src/github.com/harmony-one

## workaround for link error
#ifeq ($(shell uname),Linux)
#	ldflags=-v -extldflags "-Wl,--allow-multiple-definition"
#else
	ldflags=-v
#endif

.PHONY: setup-harmony
setup-harmony: clone-harmony build-harmony

.PHONY: clone-harmony
clone-harmony:
	mkdir -p ${harmony_root}
	cd ${harmony_root} && \
	git clone https://github.com/harmony-one/mcl.git ${GIT_OPTS} ${MCL_GIT_OPTS} && \
	git clone https://github.com/harmony-one/bls.git ${GIT_OPTS} ${BLS_GIT_OPTS} && \
	git clone https://github.com/harmony-one/harmony.git ${GIT_OPTS} ${HARMONY_GIT_OPTS}

.PHONY: build-harmony
build-harmony:
	cd ${harmony_root}/harmony && \
	go mod tidy && \
	make

.PHONY: build
build:
	. $(shell go env GOPATH)/src/github.com/harmony-one/harmony/scripts/setup_bls_build_flags.sh -v && \
	go build -ldflags='${ldflags}' -v

.PHONY: test
test:
	. $(shell go env GOPATH)/src/github.com/harmony-one/harmony/scripts/setup_bls_build_flags.sh && \
	go test -v --count=1 ./...

.PHONY: run
run:
	. $(shell go env GOPATH)/src/github.com/harmony-one/harmony/scripts/setup_bls_build_flags.sh -v && \
	go run main.go
#	. $(shell go env GOPATH)/src/github.com/harmony-one/harmony/scripts/setup_bls_build_flags.sh -v && \
	./harmony-go-sdk-sample
