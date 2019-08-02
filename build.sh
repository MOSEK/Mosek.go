#!/bin/bash

MOSEKMAJORVER=9
MOSEKMINORVER=1

case "$(uname)" in
    Linux) PLATFORM=linux64x86
        ;;
    Darwin) PLATFORM=osx64x86
        ;;
    *)  echo "Unsupported platform!"
        exit 1
        ll
esac

MOSEK_INST_DIR=${MOSEK_INST_DIR:-$HOME}
MOSEK_BIN_DIR=$MOSEK_INST_DIR/mosek/$MOSEKMAJORVER/tools/platform/$PLATFORM/bin
MOSEK_INC_DIR=$MOSEK_INST_DIR/mosek/$MOSEKMAJORVER/tools/platform/$PLATFORM/h


export CGO_CFLAGS="-I$MOSEK_INC_DIR"
export CGO_LDFLAGS="-L$MOSEK_BIN_DIR -Xlinker -rpath-link=$MOSEK_BIN_DIR"
export GOPATH=$(pwd)

#echo $CGO_CFLAGS

cd mosek
go build && \
go build examples/lo1/lo1.go && \
go build examples/qo1/qo1.go && \
go build examples/cqo1/cqo1.go
