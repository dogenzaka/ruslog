# Makefile
#

__PWD=$(shell pwd)

__gom=$(shell which gom)
__go=$(shell which go)

__goconvey=$(__PWD)/_vendor/bin/goconvey

__golint=$(__PWD)/_vendor/bin/golint


#__GOM_ENV=GOM_VENDOR_NAME=.
__GOM_ENV=

__BIN=bin
__PROG_NAME=ruslog
__PROG=$(__BIN)/$(__PROG_NAME)

__GOPATH=$(__PWD)/_vendor:$(__PWD)


all: build test info


clean:
	rm -f $(__PROG)

setup:
	$(__go) get github.com/mattn/gom

install:
	$(__GOM_ENV) $(__gom) install

install-test:
	$(__GOM_ENV) $(__gom) -test install


info:
	@echo "\nIntellJ - Go Application"
	@echo "  Run -> Environment variables"
	@echo "    Copy -> GOPATH=$(__GOPATH)"
	@echo "  Run -> Go file ✔"
	@echo "    Copy -> "$(__PWD)/src/api/main.go
	@echo "  Run -> Output executable name"
	@echo "    Copy -> $(__PROG)"
	@echo "  Run -> Arguments"
	@echo "    Copy -> -c conf/local.json"
	@echo "  Run -> Working directory"
	@echo "    Copy -> "$(__PWD)
	@echo "  Run -> Build before run -> Output directory"
	@echo "    Copy -> "$(__PWD)/bin
	@echo ""

test:
	GOPATH=$(__GOPATH) $(__go) test

lint:
	find *.go | xargs -L 1 $(__golint) -min_confidence=1.1

cover: goconvey

goconvey:
	GOPATH=$(__GOPATH) $(__goconvey) -depth=3

.PHONY: all clean setup install install-test info build build-gom test doc cover goconvey
