// Based off
// https://github.com/hashicorp/terraform/blob/master/flatmap/expand.go
package flatmap

import (
	"fmt"
	"strings"
)

func Expand(src map[string]interface{}, key string) (interface{}, error) { return mc.Expand(src, key) }
func (c *FlatMapConfig) Expand(m map[string]interface{}, key string) (interface{}, error) {
	if v, ok := m[key]; ok {
		// FIXME: Additional casting needed?
		return v, nil
	}

	// If we see ANYTHING that follows that indicates it is an array, we go with that
	arrayKey := key + c.keyDelim + "["
	for k, _ := range m {
		if strings.HasPrefix(k, arrayKey) {
			return c.expandArray(m, key)
		}
	}

	// If we see ANYTHING that indicates it is a map, we go with that
	mapKey := key + c.keyDelim
	for k, _ := range m {
		if strings.HasPrefix(k, mapKey) {
			return c.expandMap(m, key)
		}
	}

	return nil, fmt.Errorf("Key '%s' is not present in this map", key)
}

func (c *FlatMapConfig) expandArray(m map[string]interface{}, prefix string) ([]interface{}, error) {
	// prefix doesnt include training '.['
	var err error

	// Find all items in array
	var arrayKeys = make(map[string]bool)
	for k, _ := range m {
		if !strings.HasPrefix(k, prefix) {
			continue
		}
		// append string paths
		// e.g. k=.things.[0].cat.frog; prefix=.things.; afterDelim=[0].cat.frog; nextPart=[0]
		// e.g. k=.things.[0]; prefix=.; afterDelim=things.[0]; nextPart=things
		afterDelim := k[len(prefix)+1:]
		nextPart := prefix + c.keyDelim + strings.Split(afterDelim, c.keyDelim)[0]
		arrayKeys[nextPart] = true
	}

	//fmt.Println("Found arrayKeys:: ", arrayKeys)

	// FIXME: Iterate through all keys and add them in order
	// May come out out of order if some are mising
	// Option to compress arrays or leave with nil values
	var result []interface{}
	for k, _ := range arrayKeys {
		var nextItem interface{}
		nextItem, err = c.Expand(m, k)

		if err != nil {
			return result, err
		}
		result = append(result, nextItem)
	}

	return result, nil
}

func (c *FlatMapConfig) expandMap(m map[string]interface{}, prefix string) (map[string]interface{}, error) {
	// prefix does not include trailing c.keyDelim
	var err error

	var mapKeys = make(map[string]string)
	for k, _ := range m {
		if !strings.HasPrefix(k, prefix) {
			continue
		}

		// append string paths
		// e.g. k=.a.[0].cat.frog; prefix=; afterDelim=a.[0].cat.frog ; nextPart=.a
		// e.g. k=.a.[0].cat.frog; prefix=.a; afterDelim=[0].cat.frog ; nextPart=.a.[0]
		// e.g. k=.a.[0].cat.frog; prefix=.a.[0]; afterDelim=cat.frog ; nextPart=.a.[0].cat

		afterDelim := k[len(prefix)+1:]
		lastKeyComponent := strings.Split(afterDelim, c.keyDelim)[0]
		nextPart := prefix + c.keyDelim + lastKeyComponent

		//fmt.Printf("afterDelim=%s\n", afterDelim)
		//fmt.Printf("nextPart=%s\n", nextPart)

		mapKeys[nextPart] = lastKeyComponent
	}

	//fmt.Println("Found mapkeys:: ", mapKeys)

	var result = make(map[string]interface{})
	for k, lastKey := range mapKeys {
		result[lastKey], err = c.Expand(m, k)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
