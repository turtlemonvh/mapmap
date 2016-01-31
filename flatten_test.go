package mapmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	var v map[string]interface{}
	var err error

	m := map[string]interface{}{
		"cat":     "garfield",
		"dog":     "odie",
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
		".cat":                 "garfield",
		".dog":                 "odie",
		".friends.[0]":         "John",
		".turtle":              "0",
		".birdColors.cardinal": "red",
		".birdColors.blueJay":  "blue",
		".57":         int64(57),
		".doesItWork": true,
	}

	v, err = Flatten(m)
	assert.Equal(t, v, r)
	assert.Equal(t, err, nil)

}
