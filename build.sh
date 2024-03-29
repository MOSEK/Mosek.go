#!/bin/bash


cd $(dirname $0)

python3 gen/gen.py -o mosek.go gen/ais.json || exit 1

MOSEKMAJORVER=$(sed 's/\([0-9]\+\)\.\([0-9]\+\)/\1/' < MOSEKVER)
MOSEKMINORVER=$(sed 's/\([0-9]\+\)\.\([0-9]\+\)/\2/' < MOSEKVER)



case "$(uname)" in
    Linux) PLATFORM=linux64x86
        ;;
    Darwin) PLATFORM=osx64x86
        ;;
    *)  echo "Unsupported platform!"
        exit 1
esac

MOSEK_INST_DIR=${MOSEK_INST_DIR:-$HOME}
MOSEK_BIN_DIR=$MOSEK_INST_DIR/mosek/$MOSEKMAJORVER.$MOSEKMINORVER/tools/platform/$PLATFORM/bin

export CGO_LDFLAGS="-L$MOSEK_BIN_DIR -Xlinker -rpath-link=$MOSEK_BIN_DIR -Xlinker -rpath=$MOSEK_BIN_DIR"


echo CGO_CFLAGS: $CGO_CFLAGS
echo CGO_LDFLAGS: $CGO_LDFLAGS

go build && \
go build -o build/lo1 examples/lo1.go && \
go build -o build/qo1 examples/qo1.go && \
go build -o build/cqo1 examples/cqo1.go
