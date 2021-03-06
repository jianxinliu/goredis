package core

import (
	"errors"
)

// Set return ture while set a string k-v pair
// panic raise if k is a empty string
func Set(k, v string) (bool, error) {
	if len(k) == 0 {
		// panic("key cloud not be empty")
		return false, errors.New("(error) ERR wrong number of arguments for 'set' command")
	}
	Pool.strs[k] = v
	return true, nil
}

// Set k-v string with expire time, in seconds
// func Set(k,v string,ex int) bool {

// }

// Get return value with the specified key,return empty string and
// error if the specified key not exist
func Get(k string) (string, error) {
	if len(k) == 0 {
		return "", errors.New("(error) ERR wrong number of arguments for 'get' command")
	}
	v := Pool.strs[k]
	if len(v) != 0 {
		return v, nil
	} else {
		return "", errors.New("nil")
	}
}
