// Based off
// https://github.com/hashicorp/terraform/blob/master/flatmap/map.go

package flatmap

import (
	"fmt"
	"regexp"
	"strings"
)

type FlatMapConfig struct {
	keyDelim         string
	sliceKeyPattern  *regexp.Regexp
	sliceKeyTemplate func(int) string
}

var mc *FlatMapConfig

func init() {
	mc = NewFlatMapConfig()
}

func NewFlatMapConfig() *FlatMapConfig {
	mc := new(FlatMapConfig)
	mc.keyDelim = "."
	mc.sliceKeyPattern = regexp.MustCompile(`\[(\d*)\]`)
	mc.sliceKeyTemplate = func(i int) string {
		return fmt.Sprintf("[%d]", i)
	}
	return mc
}

// Can't extend map directly
// https://groups.google.com/forum/#!topic/golang-nuts/d9HMyUoKPqc
type FlatMap struct {
	Map    map[string]interface{}
	Config *FlatMapConfig
}

func NewFlatMap(m map[string]interface{}) *FlatMap {
	f := new(FlatMap)
	f.Map = m
	f.Config = mc
	return f
}

func (m *FlatMap) GetSubMap(prefix string) *FlatMap {
	var newMap = make(map[string]interface{})
	for k, v := range m.Map {
		if strings.HasPrefix(k, "."+prefix) {
			newMap[k] = v
		}
	}
	return NewFlatMap(newMap)
}

func (m *FlatMap) Reshuffle(oldpath string, newpath string) error {
	var tomove = make(map[string]interface{})
	for k, v := range m.Map {
		if strings.HasPrefix(k, "."+oldpath) {
			tomove[k] = v
		}
	}

	for k, v := range tomove {
		delete(m.Map, k)

		// k=.a.b.c ; oldpath=a.b; newpath=c
		newKey := newpath + k[len(oldpath)+1:]
		m.Map["."+newKey] = v
	}

	return nil
}

// Contains returns true if the map contains the given key.
func (m *FlatMap) Contains(key string) bool {
	for _, k := range m.Keys() {
		if k == key {
			return true
		}
	}
	return false
}

// Delete deletes a key out of the map with the given prefix.
func (m *FlatMap) Delete(prefix string) {
	prefix = m.Config.keyDelim + prefix
	for k, _ := range m.Map {
		match := k == prefix
		if !match {
			if !strings.HasPrefix(k, prefix) {
				continue
			}

			if k[len(prefix):len(prefix)+1] != m.Config.keyDelim {
				continue
			}
		}

		delete(m.Map, k)
	}
}

// Keys returns list of top-level keys.
func (m *FlatMap) Keys() []string {
	var mapKeys = make(map[string]bool)
	for k, _ := range m.Map {
		mapKeys[strings.Split(k[1:], m.Config.keyDelim)[0]] = true
	}

	var result []string
	for k, _ := range mapKeys {
		result = append(result, k)
	}
	return result
}

// Merge merges the contents of the other Map into this one.
//
// Any shared top level keys will be overwritten.
func (m *FlatMap) Merge(m2 *FlatMap) {
	for _, prefix := range m2.Keys() {
		prefix = m.Config.keyDelim + prefix
		m.Delete(prefix[1:])

		for k, v := range m2.Map {
			if strings.HasPrefix(k, prefix) {
				m.Map[k] = v
			}
		}
	}
}
