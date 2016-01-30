package mapmap

import (
	_ "github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
FIXME:
- add more diabolical tests
- add tests for set that conflict with each other (one creates a slice, the next creates a slice)
    - add function to test for this

*/

func TestFlatten(t *testing.T) {
	var v map[string]interface{}
	var err error

	m := map[string]interface{}{
		"cat":     "garfield",
		"dog":     "oddie",
		"friends": []interface{}{"John"},
		"turtle":  "0",
		"birdColors": map[string]interface{}{
			"cardinal": "red",
			"blueJay":  "blue",
		},
		"57":         57,
		"doesItWork": true,
	}

	r := map[string]interface{}{
		"cat":                 "garfield",
		"dog":                 "oddie",
		"friends.[0]":         "John",
		"turtle":              "0",
		"birdColors.cardinal": "red",
		"birdColors.blueJay":  "blue",
		"57":                  int64(57),
		"doesItWork":          true,
	}
	//err = SetNested("cat", "milo", &m, ".")
	//assert.Equal(t, err, nil)
	v, err = Flatten(m)
	assert.Equal(t, v, r)
	assert.Equal(t, err, nil)

}
