Mapmap [![GoDoc](https://godoc.org/github.com/turtlemonvh/mapmap?status.svg)](https://godoc.org/github.com/turtlemonvh/mapmap)
===

> It maps your maps!

> **WORK IN PROGRESS:** Doesn't actually do anything useful yet

## Quick Start

Mapmap is for doing quick transformations on `map[string]interface{}` and `[]interface{}` objects in golang (the stuff you get back when parsing json).

Create a slice of `Mapper`s, and then call `mapmap.Mapit` on your object.  The result is a new object with fields reshuffled and transformed.

## Testing

    go test ./...

## Credits

* Built with the awesome [`cast` library](https://github.com/spf13/cast) from spf13. 
* Inspired by the [`viper` library](https://github.com/spf13/cast) from spf13.
* Lots of source taken from the [`flatmap` library](https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go) in terraform.
    * Mostly I just manage maps with interfaces and return errors instead of panicing.

## TODO

* add transform functions
* add validators
* add type handling
* add error handling (returning array of errors) in addition to standard handling
