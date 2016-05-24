SHELL=/bin/bash
CMD=locofo
GOBUILD=go build
USER=stvnwrgs
REPOSITORY=locofo
VERSION=$(shell git describe --always --tags --dirty | cut -f1 -d"-")

XC_ARCH=${XC_ARCH:-"386 amd64 arm"}
XC_OS=${XC_OS:-linux darwin freebsd openbsd solaris}

build:
	XC_OS=${XC_OS} XC_ARCH=${XC_ARCH} CMD=${CMD} ./scripts/build.sh
release-gh:
	USER=${USER} REPOSITORY=${REPOSITORY} VERSION=${VERSION} CMD=${CMD} ./scripts/release.sh

