// This package contains misc tools functions
package tools

import (
	"fmt"
	"strings"
)

// This tool function takes a map and delete all the fields that are absent from `wanted` []string
func DeleteUnwanted(_map map[string]interface{}, wanted []string, oldPath string) {
	for key, value := range _map {
		var path string

		if oldPath != "" {
			path = oldPath + "." + key
		} else {
			path = key
		}

		if innerMap, ok := value.(map[string]interface{}); ok {
			DeleteUnwanted(innerMap, wanted, path)
		} else {
			found := false
			for i, length := 0, len(wanted); i < length; i++ {
				if path == wanted[i] {
					found = true
					break
				}
			}
			if !found {
				delete(_map, key)
			}
		}
	}
}

// This tool function accepts a dotted path (like "grandparent.parent.child") and browse the provided map for the matching value
func ReadDeep(_map map[string]interface{}, path string) (interface{}, error) {
	indexes := strings.Split(path, ".")
	for i := 0; i < len(indexes)-1; i++ {
		index := indexes[i]
		// if one element of the path is nil, the return value does not exist
		if _map[index] == nil {
			return nil, nil
		} else {
			// the value exists, but is it a map?
			if m, ok := _map[index].(map[string]interface{}); !ok {
				return nil, fmt.Errorf("Invalid path: %s path element is not a map", index)
			} else {
				_map = m
			}
		}
	}
	return _map[indexes[len(indexes)-1]], nil
}

// This tool function accepts a dotted path (like "grandparent.parent.child") and write the provided value in the map container, creating inner maps if needed
func WriteDeep(_map map[string]interface{}, path string, value interface{}) error {
	indexes := strings.Split(path, ".")
	for i := 0; i < len(indexes); i++ {
		// alias
		index := indexes[i]

		// if last index, write the value
		if i == len(indexes)-1 {
			_map[index] = value
		} else {
			// creating inner map if not exists
			if _map[index] == nil {
				_map[index] = make(map[string]interface{})
			}
			// type checking
			if m, ok := _map[index].(map[string]interface{}); !ok {
				return fmt.Errorf("Invalid path: % path element already exists and is not a map", index)
			} else {
				// one map deeper!
				_map = m
			}
		}
	}
	return nil
}
