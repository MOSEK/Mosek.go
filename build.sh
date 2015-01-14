#!/bin/bash


if [ -z "$1" ] ; then
  echo "Examples:"
  echo "  $1 install mosek"
  echo "  $1 install examples/..."
  echo "" 
  echo "NOTE: The environment variable DISTBASE must be set and point to the platform"
  echo "      directory of a valid MOSEK installation (e.g. .../platform/linux64x86)."
  exit
fi    

if [ -z "$DISTBASE" ]; then
  echo "define DISTBASE to point to the distribution directory (e.g. linux64x86)"
  exit 1
fi


cd $(dirname $0)
CGO_CFLAGS=-I$DISTBASE/h \
CGO_LDFLAGS=-L$DISTBASE/bin \
LD_LIBRARY_PATH=$DISTBASE/bin \
GOPATH=$(pwd) \
  go "$@"
