# *DISCLAIMER*

This software is not officially supported. We (MOSEK) accept bug reports and
may fix errors and bugs, but provide no guarantee on how or how quickly we do so.

The API has been tested on Linux x86-64, but may work on other posix platforms
as well. It may require some tweaking to get it to build on Windows.

# Versioning

The `Major.Minor` part of the version corresponds to the MOSEK Version it requires.

# Requirements

- MOSEK tools
- Go v1.21+

The MOSEK Go interface does not depend on any third-party packages, only Go standard libraries. 

# Go interface for MOSEK

This is a thin Go interface for MOSEK.

Currently implemented:

- Nearly all relevant functions (and a lot of more or less useless ones) are mapped
- Stream printing callbacks
- Progress and information callbacks

It it not thoroughly tested.

## Documentation

Currently, it is recommened to use `go doc ...` and the official C documentation [official C API documentation](https://docs.mosek.com/9.2/capi/index.html).

# Using MOSEK from another project

When building a project that uses MOSEK, MOSEK must be available (the specific
version is defined in the file `MOSEKVER`). Usually MOSEK will not be
globally available, so locations the library library must be
specified when building a project that imports MOSEK, e.g.
```
CGO_LDFLAGS="-L$MOSEKBINDIR" go build
```
or alternatively
```
LPATH=$MOSEKBINDIR go build
```

# Building examples

To build the included examples do 
```
CGO_LDFLAGS="-L$MOSEKBINDIR" go build examples/lo1.go
CGO_LDFLAGS="-L$MOSEKBINDIR" go build examples/qo1.go
CGO_LDFLAGS="-L$MOSEKBINDIR" go build examples/cqo1.go
```
