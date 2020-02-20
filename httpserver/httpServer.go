package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/jianxin/goredis/client/rpcClient"
	// rpcServer "github.com/jianxin/goredis/rpcserver"
)

func main() {
	// 控制 rpcClient 的创建
	// rpcServer.UnDialed = !rpcServer.UnDialed
	client.Init()
	http.HandleFunc("/exec", exec)
	for true {
		fmt.Println("httpserver listening on 127.0.0.1:9090 ....")
		http.Handle("/", http.FileServer(http.Dir("./static")))
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}
}

type request struct {
	cmd string
}

func exec(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cmd := r.Form["cmd"][0]
	if cmd == "exit" {
		fmt.Fprintf(w, "Bye!")
		os.Exit(0)
	}
	// fmt.Println(cmd)
	cmdType := client.FirstWord(cmd)
	// fmt.Println(cmdType + "  " + cmd)
	if client.RpcServices[cmdType] == true {
		rep, err := client.DialRPC(cmdType, cmd)
		if err != nil {
			fmt.Fprintf(w, "nil")
		} else {
			fmt.Fprintf(w, rep)
		}
	} else {
		fmt.Fprintf(w, "unsupported command!")
	}
}
