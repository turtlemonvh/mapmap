package flatmap

import (
	_ "fmt"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ListContains(l []string, i string) bool {
	for _, li := range l {
		if li == i {
			return true
		}
	}
	return false
}

func ListMatch(l1 []string, l2 []string) bool {
	for _, l1i := range l1 {
		if !ListContains(l2, l1i) {
			return false
		}
	}
	return true
}

func TestSubMap(t *testing.T) {
	m := map[string]interface{}{
		".cat":                 "garfield",
		".dog":                 "odie",
		".birdColors.cardinal": "red",
		".birdColors.blueJay":  "blue",
		".57": int64(57),
	}

	f := NewFlatMap(m)
	assert.Equal(t, len(f.Keys()), 4)
	fsub := f.GetSubMap("birdColors")
	assert.Equal(t, len(fsub.Keys()), 1)
	assert.True(t, ListMatch(fsub.Keys(), []string{"birdColors"}))
}

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

	// FIXME: Buggy because of array order issues
	assert.Equal(t, cast.ToStringMap(v), r2)
	assert.Equal(t, err, nil)
}
