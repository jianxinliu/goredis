package rpcserver

import (
	"errors"
	"fmt"
	"github.com/jianxin/goredis/core"
	"log"
	"net"
	"net/rpc"
	// "strings"
)

type CmderService struct{}

func (cmdService *CmderService) Set(req string, reply *string) error {
	// k,v,params := splitCmd(req)
	fmt.Println(req)
	k, v := "name", "jianxin"
	done := core.Set(k, v)
	if done {
		*reply = "OK"
		return nil
	} else {
		*reply = ""
		return errors.New(fmt.Sprintf("set with key:%s,value:%s failed", k, v))
	}
}

// cmd like "SET name jianxin 60"
// identify wrong command wtih panic
// func splitCmd(cmd string) (k,v string,params []string){
// 	filds := strings.Fields(cmd)
// 	return "","",[]string{""}
// }

// start the goredis server
func Start() {
	rpc.RegisterName("CmderService", new(CmderService))

	listener, err := net.Listen("tcp", ":8064")
	checkErr(err)
	conn, err := listener.Accept()
	checkErr(err)
	rpc.ServeConn(conn)
}

// func router(cmdStr string,handler func(cmd string) (string,error)) (string,error) {

// }

func checkErr(err error) {
	if err != nil {
		log.Fatal("Accpter err:", err)
	}
}
