# Mapmap

> It maps your maps!

> **WORK IN PROGRESS:** Doesn't actually do anything useful yet

## Quick Start

Mapmap is for doing quick transformations on `map[string]interface{}` and `[]interface{}` objects in golang (the stuff you get back when parsing json).

Create a slice of `Mapper`s, and then call `mapmap.Mapit` on your object.  The result is a new object with fields reshuffled and transformed.

## Credits

* Built with the awesome [`cast` library](https://github.com/spf13/cast) from spf13. 
* Inspired by the [`viper` library](https://github.com/spf13/cast) from spf13.
* Lots of source taken from the [`flatmap` library](https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go) in terraform.
    * Mostly I just manage maps with interfaces and return errors instead of panicing.

## TODO

* actual mapmap implementation
* escape key delimiter when flattening, unescape when expanding
* tests for changing key delimiter
    * e.g. merging maps with different delimiters
* tougher tests in general
