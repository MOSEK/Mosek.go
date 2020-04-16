# *DISCLAIMER*

This software is experimental. We (MOSEK) may fix errors and bugs, but provide
no guarantee on how or how quickly we do so.

The API has been tested on 64bit Linux, but may work on other posix platforms as well.


# Requirements

- MOSEK tools
- Go v1.12+

The MOSEK Go interface does not depend on any third-party packages.

# Go interface for MOSEK

This is a thin Go interface for MOSEK.

Currently implemented:

- Nearly all relevant functions (and a lot of more or less useless ones) are mapped
- Stream printing callbacks
- Progress and information callbacks

## Documentation

The documentation contains an API reference covering all functions and
constants in the interface. It may contain dead links since many sections are
not included in the manual yet.

- [API reference](docs/mosek.rst)
- [Constants reference](docs/mosek_const.rst).

For a more complete version of the documentation, see the [official C API documentation](https://docs.mosek.com/9.2/capi/index.html).

# Using MOSEK from another project

You can use Mosek from a go project by importing it by its full name:

```
import mosek "github.com/mosek/mosek.go"
```

When building using `go build`, MOSEK must be available (the specific
version is defined in the file `MOSEKVER`). Usually MOSEK will not be
globally available, so locations of header file and library must be
specified when building a project that imports MOSEK:

```
export CGO_CFLAGS="-I$HOME/mosek/9.2/tools/platform/linux64x86/h"
export CGO_LDFLAGS="-L$HOME/mosek/9.2/tools/platform/linux64x86/bin"
go build
```

# Building examples

Run the included script `build.sh` to build examples.
