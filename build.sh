#!/bin/bash


cd $(dirname $0)

python3 gen/gen.py -o functions.go gen/ais.json || exit 1

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
go run examples/lo1.go && \
go run examples/qo1.go && \
go run examples/cqo1.go && \
go run examples/acc1.go && \
go run examples/acc2.go && \
go run examples/djc1.go && \
go run examples/sdo1.go && \
go run examples/pow1.go && \
go run examples/callback.go && \
go run examples/portfolio_1_basic.go && \
go run examples/portfolio_2_frontier.go && \
go run examples/portfolio_3_impact.go && \
go run examples/concurrent1.go examples/25fv47.mps
