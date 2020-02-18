package main

import (
	// "github.com/jianxin/goredis/core"
	"github.com/jianxin/goredis/rpcserver"
)

func main() {
	// core.Run()
	rpcserver.Start()
}
