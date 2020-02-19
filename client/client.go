package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

const host string = "127.0.0.1:8064"

var rpcServices = map[string]bool{
	"Set":  true,
	"Get":  true,
	"Keys": true,
}

var client *rpc.Client

func init() {
	cli, err := rpc.Dial("tcp", host)
	if err != nil {
		log.Fatal("Dial err:", err)
	}
	client = cli
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(host + "> ")
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "exit" {
			os.Exit(0)
		}
		cmdType := firstWord(cmd)
		if rpcServices[cmdType] == true {
			rep, err := dialRPC(cmdType, cmd)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(rep)
			}
			fmt.Print(host + "> ")
		} else {
			fmt.Println("unsupported command!")
			fmt.Print(host + "> ")
		}
	}
}

func firstWord(s string) string {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
	}
	return toUpperCamelcase(s[0:i])
}

func toUpperCamelcase(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:len(s)]
}

func dialRPC(cmdType string, cmd string) (string, error) {
	var reply string
	err := client.Call("CmderService."+cmdType, cmd, &reply) // 命名空间
	if err != nil {
		return reply, err
	} else {
		return reply, nil
	}
}
