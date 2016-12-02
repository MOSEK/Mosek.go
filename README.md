# *DISCLAIMER*

This software is experimental. We (MOSEK) may fix errors and bugs, but provide
no guarantee on how or how quickly we do so.

The API has been tested on 64bit Linux, but may work on other posix platforms as well.

# Go interface for MOSEK

This is a thin Go interface for MOSEK.

Currently implemented:

- Nearly all relevant functions (and a lot of more or less useless ones) are mapped
- Stream printing callbacks
- Progress and information callbacks

What will probably never be supported:

- General convex API

## Documentation

The documentation contains an API reference covering all functions and
constants in the interface. It may contain dead links since many sections are
not included in the manual yet. 

- [API reference](docs/mosek.rst)
- [Constants reference](docs/mosek_const.rst).

# Building

Building requires MOSEK 8 and `go` 1.2.1 or later. By default the
builder script will look in `$HOME` for a mosek installation; this can be overridden by setting the `MOSEK_INST_DIR` variable:
```
MOSEK_INST_DIR=$HOME/local build.sh
```
This will build the library and all examples.
