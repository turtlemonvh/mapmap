package mapmap

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"regexp"
	"strings"
)

/*

We want a map of strings to functions that can

Get function:
https://github.com/spf13/viper/blob/master/viper.go#L457

*/

var SliceKeyPattern = regexp.MustCompile(`\[(\d*)\]`)

type Mapper struct {
	input      string // dot delimited path to input field in map
	output     string // dot delimited path to output field in map
	typeString string // string representation of type
	checkType  string
	keyDelim   string
}

func (m *Mapper) Map(inMap *map[string]interface{}, outMap *map[string]interface{}) error {
	return nil
}

// Only works with map[string]interface{} and []interface{}
func GetNested(key string, src interface{}, keyDelim string) (val interface{}, err error) {
	var isNested bool
	var keyParts []string
	if strings.Contains(key, keyDelim) {
		isNested = true
		keyParts = strings.Split(key, keyDelim)
	} else {
		isNested = false
		keyParts = []string{key}
	}
	currentKey := keyParts[0]

	fmt.Printf("currentKey: %s ; keyParts: %s, src: %s\n", currentKey, keyParts, src)

	// Handle src as either map or slice
	if reflect.TypeOf(src).Kind() == reflect.Map {
		fmt.Println("type: map")
		s, errp := cast.ToStringMapE(src)
		fmt.Println("cast: ", s, errp)
		val = cast.ToStringMap(src)[currentKey]
		fmt.Println("val: ", val)
	} else if reflect.TypeOf(src).Kind() == reflect.Slice {
		fmt.Println("type: slice")
		// Extract integer component from next key
		var intKey int
		ms := SliceKeyPattern.FindStringSubmatch(currentKey)
		if len(ms) < 2 {
			return nil, fmt.Errorf("The key '%s' cannot be used to index into slice '%s'", currentKey, src)
		}
		intKey, err = cast.ToIntE(ms[1])
		if err != nil {
			return nil, err
		}

		fmt.Printf("before: %s, after: %d\n", currentKey, intKey)

		subSrc := cast.ToSlice(src)
		if intKey >= 0 && intKey < len(subSrc) {
			val = subSrc[intKey]
		} else {
			val = nil
		}
	} else {
		fmt.Println("type: ???")
		// Tried to index into something that's not an array or map
		val = nil
	}

	fmt.Printf("value: %s \n", val)

	if val == nil {
		return val, err
	}

	if isNested {
		subPath := strings.Join(keyParts[1:], keyDelim)
		val, err = GetNested(subPath, val, keyDelim)
	}

	return val, err
}

func MapIt(inMap *map[string]interface{}, outMap *map[string]interface{}, mappers []Mapper, exitEarly bool) []error {
	return []error{nil}
}
