package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8064")
	if err != nil {
		log.Fatal("Dial err:", err)
	}
	var reply string
	err = client.Call("CmderService.Set", "SET name jianxin", &reply) // 命名空间
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
