package core

import (
// "errors"
// . "fmt"
)

const (
	SET = iota
	GET
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
