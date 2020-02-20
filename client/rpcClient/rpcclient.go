package rpcClient

import (
	"log"
	"net/rpc"
	"strings"
	// "fmt"
)

var RpcServices = map[string]bool{
	"Set":  true,
	"Get":  true,
	"Keys": true,
}

const Host string = "127.0.0.1:8064"

var client *rpc.Client

func Init() {
	cli, err := rpc.Dial("tcp", Host)
	if err != nil {
		log.Fatal("Dial err:", err)
	}
	client = cli
}

func DialRPC(cmdType string, cmd string) (string, error) {
	var reply string
	// fmt.Println(cmdType + "****" + cmd)
	err := client.Call("CmderService."+cmdType, cmd, &reply) // 命名空间
	if err != nil {
		return reply, err
	} else {
		return reply, nil
	}
}

func FirstWord(s string) string {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
	}
	return toUpperCamelcase(s[0:i])
}

func toUpperCamelcase(s string) string {
	s = strings.ToLower(s)
	return strings.ToUpper(s[0:1]) + s[1:len(s)]
}
