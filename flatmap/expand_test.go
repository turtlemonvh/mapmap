package flatmap

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpand(t *testing.T) {
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

	v, err = Expand(m, ".cat")
	assert.Equal(t, cast.ToString(v), "garfield")
	assert.Equal(t, err, nil)

	r := map[string]interface{}{
		"cat":     "garfield",
		"dog":     "odie",
		"friends": []interface{}{"John"},
		"turtle":  "0",
		"birdColors": map[string]interface{}{
			"cardinal": "red",
			"blueJay":  "blue",
		},
		"57":         int64(57),
		"doesItWork": true,
	}

	v, err = Expand(m, "")
	assert.Equal(t, cast.ToStringMap(v), r)
	assert.Equal(t, err, nil)
}
