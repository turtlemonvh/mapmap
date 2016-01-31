package mapmap

import (
	"fmt"
	_ "github.com/spf13/cast"
	_ "reflect"
	"regexp"
	_ "strings"
)

/*

We want a map of strings to functions that can

Get function:
https://github.com/spf13/viper/blob/master/viper.go#L457

https://github.com/hashicorp/terraform/tree/master/flatmap
https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go

*/

type MapperConfig struct {
	keyDelim         string
	sliceKeyPattern  *regexp.Regexp
	sliceKeyTemplate func(int) string
}

var mc *MapperConfig

func init() {
	mc = New()
}

func New() *MapperConfig {
	mc := new(MapperConfig)
	mc.keyDelim = "."
	mc.sliceKeyPattern = regexp.MustCompile(`\[(\d*)\]`)
	mc.sliceKeyTemplate = func(i int) string {
		return fmt.Sprintf("[%d]", i)
	}
	return mc
}

type Mapper struct {
	input      string // dot delimited path to input field in map
	output     string // dot delimited path to output field in map
	typeString string // string representation of type
	checkType  string
	exitEarly  bool
}

func (m *Mapper) Map(inMap *map[string]interface{}, outMap *map[string]interface{}) error {
	return nil
}

func MapIt(inMap interface{}, mappers []Mapper) (interface{}, []error) {
	return nil, []error{nil}
}
