package mapmap

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapMap(t *testing.T) {
	var v interface{}
	var processingErrors []error
	var err error

	m := map[string]interface{}{
		"cat":        "garfield",
		"dog":        "odie",
		"friends":    []interface{}{"John"},
		"turtle":     "0",
		"57":         int64(57),
		"doesItWork": true,
	}

	r := map[string]interface{}{
		"cat":          "garfield",
		"frog":         "0",
		"myOnlyFriend": "John",
	}

	var mappers []*Mapper
	mappers = append(mappers, NewMapper("cat", "cat"))
	mappers = append(mappers, NewMapper("turtle", "frog"))
	mappers = append(mappers, NewMapper("friends.[0]", "myOnlyFriend"))

	v, processingErrors, err = MapIt(m, mappers)

	assert.Equal(t, cast.ToStringMap(v), r)
	assert.Equal(t, processingErrors, []error{})
	assert.Equal(t, len(processingErrors), 0)
	assert.Equal(t, err, nil)
}
