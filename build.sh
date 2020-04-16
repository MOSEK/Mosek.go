#!/bin/bash


cd $(dirname $0)


MOSEKMAJORVER=$(sed 's/\([0-9]\+\)\.\([0-9]\+\)/\1/' < MOSEKVER)
MOSEKMINORVER=$(sed 's/\([0-9]\+\)\.\([0-9]\+\)/\2/' < MOSEKVER)



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
MOSEK_BIN_DIR=$MOSEK_INST_DIR/mosek/$MOSEKMAJORVER.$MOSEKMINORVER/tools/platform/$PLATFORM/bin
MOSEK_INC_DIR=$MOSEK_INST_DIR/mosek/$MOSEKMAJORVER.$MOSEKMINORVER/tools/platform/$PLATFORM/h


export CGO_CFLAGS="-I$MOSEK_INC_DIR"
export CGO_LDFLAGS="-L$MOSEK_BIN_DIR -Xlinker -rpath-link=$MOSEK_BIN_DIR"

go build && \
go install examples/lo1.go && \
go install examples/qo1.go && \
go install examples/cqo1.go
