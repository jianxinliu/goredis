package core

import (
	"errors"
	// . "fmt"
	"strings"
)

type CmdType struct{}

// goredis storage
type memPool struct {
	strs map[string]string
	// lists
	// sets
}

var Pool memPool

func init() {
	Pool = memPool{make(map[string]string)}
}

// need improve!
// Keys return keys like the pattern
func Keys(pattern string) ([]string, error) {
	keys := make([]string, 0, len(Pool.strs))
	if pattern == "*" {
		for k := range Pool.strs {
			keys = append(keys, k)
		}
	} else {
		for k := range Pool.strs {
			if strings.Index(k, pattern) != -1 {
				keys = append(keys, k)
			}
		}
	}
	if len(keys) == 0 {
		return keys, errors.New("(empty list or set)")
	} else {
		return keys, nil
	}
}
