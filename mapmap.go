package mapmap

import (
	_ "fmt"
	_ "github.com/spf13/cast"
	"github.com/turtlemonvh/mapmap/flatmap"
	_ "strings"
)

/*

We want a map of strings to functions that can

Get function:
https://github.com/spf13/viper/blob/master/viper.go#L457

https://github.com/hashicorp/terraform/tree/master/flatmap
https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go

*/

type Mapper struct {
	InputPath  string // dot delimited path to input field in map
	OutputPath string // dot delimited path to output field in map
	typeString string // string representation of type
	checkType  string
	exitEarly  bool
}

func NewMapper(input string, output string) *Mapper {
	m := new(Mapper)
	m.InputPath = input
	m.OutputPath = output
	return m
}

func MapIt(inMap interface{}, mappers []*Mapper) (interface{}, []error) {
	inMapFlat, err := flatmap.Flatten(inMap)
	if err != nil {
		return nil, []error{err}
	}

	in := flatmap.NewFlatMap(inMapFlat)
	out := flatmap.NewFlatMap(make(map[string]interface{}))

	// Just reshuffle for now
	for _, m := range mappers {
		tf := in.GetSubMap(m.InputPath)
		tf.Reshuffle(m.InputPath, m.OutputPath)
		out.Merge(tf)
	}
	exp, err := flatmap.Expand(out.Map, "")

	return exp, []error{}
}
