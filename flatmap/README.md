Flatmap [![GoDoc](https://godoc.org/github.com/turtlemonvh/mapmap/flatmap?status.svg)](https://godoc.org/github.com/turtlemonvh/mapmap/flatmap)
===

> **WORK IN PROGRESS:** Works, but not as well as it looks like it does :P

Flatten out and reexpand `map[string]interface{}` and `[]interface{}` objects.

Aims to be configurable.

## TODO

* escape key delimiter when flattening, unescape when expanding
* tests for changing key delimiter
    * e.g. merging maps with different delimiters
* tougher tests in general