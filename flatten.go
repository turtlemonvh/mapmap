// Based off:
// https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go

package mapmap

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

func Flatten(src interface{}) (map[string]interface{}, error) { return mc.Flatten(src) }
func (c *MapperConfig) Flatten(src interface{}) (map[string]interface{}, error) {
	var err error

	result := make(map[string]interface{})

	if reflect.TypeOf(src).Kind() == reflect.Map {
		var s map[string]interface{}
		s, err = cast.ToStringMapE(src)
		if err != nil {
			return result, err
		}

		for k, raw := range s {
			err = c.flatten(result, k, reflect.ValueOf(raw))
			if err != nil {
				break
			}
		}

	} else if reflect.TypeOf(src).Kind() == reflect.Slice {
		var s []interface{}
		s, err = cast.ToSliceE(src)
		if err != nil {
			return result, err
		}

		for i, raw := range s {
			err = c.flatten(result, c.sliceKeyTemplate(i), reflect.ValueOf(raw))
			if err != nil {
				break
			}
		}
	} else {
		err = fmt.Errorf("Unknown type for interface")
	}

	return result, err
}

func (c *MapperConfig) flatten(result map[string]interface{}, prefix string, v reflect.Value) error {
	var err error

	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	// Set as type interface
	switch v.Kind() {
	case reflect.Bool:
		result[c.keyDelim+prefix] = v.Bool()
	case reflect.Int:
		result[c.keyDelim+prefix] = v.Int()
	case reflect.String:
		result[c.keyDelim+prefix] = v.String()
	case reflect.Map:
		err = c.flattenMap(result, prefix, v)
	case reflect.Slice:
		err = c.flattenSlice(result, prefix, v)
	default:
		err = fmt.Errorf("Unknown primitive type found for value: '%q'", v)
	}

	return err
}

func (c *MapperConfig) flattenMap(result map[string]interface{}, prefix string, v reflect.Value) error {
	var err error

	for _, k := range v.MapKeys() {
		if k.Kind() == reflect.Interface {
			k = k.Elem()
		}

		if k.Kind() != reflect.String {
			err = fmt.Errorf("%s: map key is not string: %s", prefix, k)
			break
		}

		err = c.flatten(result, fmt.Sprintf("%s%s%s", prefix, c.keyDelim, k.String()), v.MapIndex(k))
		if err != nil {
			break
		}
	}

	return err
}

func (c *MapperConfig) flattenSlice(result map[string]interface{}, prefix string, v reflect.Value) error {
	var err error

	for i := 0; i < v.Len(); i++ {
		err = c.flatten(result, fmt.Sprintf("%s%s%s", prefix, c.keyDelim, c.sliceKeyTemplate(i)), v.Index(i))
		if err != nil {
			break
		}
	}

	return err
}
