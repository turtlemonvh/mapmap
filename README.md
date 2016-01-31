Mapmap [![GoDoc](https://godoc.org/github.com/turtlemonvh/mapmap?status.svg)](https://godoc.org/github.com/turtlemonvh/mapmap)
===

> It maps your maps!

> **WORK IN PROGRESS:** Very limited features so far.

## Quick Start

Mapmap is for doing quick transformations on `map[string]interface{}` and `[]interface{}` objects in golang (the stuff you get back when parsing json).

Basic usage for shuffling around fields is as follows:

```go
package main

import (
    "fmt"

    "github.com/turtlemonvh/mapmap"
)

func main() {

    // You need an object to map
    m := map[string]interface{}{
        "cat":        "garfield",
        "dog":        "odie",
        "friends":    []interface{}{"John"},
        "turtle":     "0",
        "57":         int64(57),
        "doesItWork": true,
    }

    // Create a slice of `Mapper`s describing transformations to run
    var mappers []*mapmap.Mapper
    mappers = append(mappers, mapmap.NewMapper("cat", "cat"))
    mappers = append(mappers, mapmap.NewMapper("turtle", "frog"))
    mappers = append(mappers, mapmap.NewMapper("friends.[0]", "myOnlyFriend"))

    // Run the tranformations
    new_map, processingErrors, err = mapmap.MapIt(m, mappers)

    fmt.Println(new_map)
}
```

This would result in the following value for `new_map`:

```go
map[string]interface{}{
    "cat":          "garfield",
    "frog":         "0",
    "myOnlyFriend": "John",
}
```

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
