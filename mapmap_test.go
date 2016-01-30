package mapmap

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNestedFromMap(t *testing.T) {
	var v interface{}
	var err error

	m := map[string]interface{}{
		"cat":     "garfield",
		"dog":     "oddie",
		"friends": []interface{}{"John"},
	}
	v, err = GetNested("cat", m, ".")
	assert.Equal(t, cast.ToString(v), "garfield")
	assert.Equal(t, err, nil)

	v, err = GetNested("friends", m, ".")
	assert.Equal(t, cast.ToStringSlice(v), []string{"John"})
	assert.Equal(t, err, nil)

	v, err = GetNested("friends.[0]", m, ".")
	assert.Equal(t, cast.ToString(v), "John")
	assert.Equal(t, err, nil)
}

func TestGetNestedFromSlice(t *testing.T) {
	var v interface{}
	var err error

	m := []interface{}{
		"cat",
		"dog",
		[]interface{}{"John"},
	}
	v, err = GetNested("[0]", m, ".")
	assert.Equal(t, cast.ToString(v), "cat")
	assert.Equal(t, err, nil)

	v, err = GetNested("[2]", m, ".")
	assert.Equal(t, cast.ToStringSlice(v), []string{"John"})
	assert.Equal(t, err, nil)

	v, err = GetNested("[2].[0]", m, ".")
	assert.Equal(t, cast.ToString(v), "John")
	assert.Equal(t, err, nil)
}
