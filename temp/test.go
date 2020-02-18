package main

import (
	. "fmt"
	"github.com/jianxin/goredis/core"
)

func main() {
	strs := make(map[string]string, 3)
	strs["name"] = "jianxin"
	Println(len(strs["name"]))
	Println(len(strs["age"]))

	core.Set("name", "jianxin")

	v, err := core.Get("name")
	if err != nil {
		Println(err)
	}
	Println(v)
}
