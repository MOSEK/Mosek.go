#!/bin/bash

if [ -z "$DISTBASE" ]; then
  echo "define DISTBASE to point to the distribution directory (e.g. linux64x86)"
  exit 1
  #DISTBASE=$HOME/mosek/7/tools/platform/linux64x86
fi

cd $(dirname $0)
export CGO_CFLAGS=-I$DISTBASE/h
export CGO_LDFLAGS=-L$DISTBASE/bin
export LD_LIBRARY_PATH=$DISTBASE/bin
export GOPATH=$(pwd)

go "$@"
