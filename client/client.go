package main

import (
	"bufio"
	"fmt"
	"os"

	client "github.com/jianxin/goredis/client/rpcClient"
)

func main() {
	client.Init()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(client.Host + "> ")
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "exit" {
			os.Exit(0)
		}
		cmdType := client.FirstWord(cmd)
		// fmt.Println(cmdType + "  " + cmd)
		if client.RpcServices[cmdType] == true {
			rep, err := client.DialRPC(cmdType, cmd)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(rep)
			}
			fmt.Print(client.Host + "> ")
		} else {
			fmt.Println("unsupported command!")
			fmt.Print(client.Host + "> ")
		}
	}
}
