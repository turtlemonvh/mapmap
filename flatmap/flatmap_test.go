package flatmap

import (
	_ "fmt"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatMap(t *testing.T) {
	var v interface{}
	var err error

	m := map[string]interface{}{
		".cat":                 "garfield",
		".dog":                 "odie",
		".friends.[0]":         "John",
		".turtle":              "0",
		".birdColors.cardinal": "red",
		".birdColors.blueJay":  "blue",
		".57":         int64(57),
		".doesItWork": true,
	}

	f := NewFlatMap(m)
	assert.Equal(t, f.Contains("cat"), true)
	assert.Equal(t, f.Contains(".cat"), false)
	assert.Equal(t, f.Contains(".birdColors.cardinal"), false)

	f.Delete("birdColors")

	r := map[string]interface{}{
		"cat":        "garfield",
		"dog":        "odie",
		"friends":    []interface{}{"John"},
		"turtle":     "0",
		"57":         int64(57),
		"doesItWork": true,
	}

	v, err = Expand(f.Map, "")
	assert.Equal(t, cast.ToStringMap(v), r)
	assert.Equal(t, err, nil)

	m2 := map[string]interface{}{
		".friends.[0]": "Odie",
		".friends.[1]": "John",
	}

	f2 := NewFlatMap(m2)
	f.Merge(f2)

	r2 := map[string]interface{}{
		"cat":        "garfield",
		"dog":        "odie",
		"friends":    []interface{}{"Odie", "John"},
		"turtle":     "0",
		"57":         int64(57),
		"doesItWork": true,
	}

	v, err = Expand(f.Map, "")
	assert.Equal(t, cast.ToStringMap(v), r2)
	assert.Equal(t, err, nil)
}
