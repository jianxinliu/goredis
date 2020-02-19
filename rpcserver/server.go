package rpcserver

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"strings"

	"github.com/jianxin/goredis/core"
)

type CmderService struct{}

func (c CmderService) Set(req string, reply *string) error {
	// log.Println("Set been called...")
	fields := strings.Fields(req)
	if len(fields) < 3 {
		return errors.New("(error) ERR wrong number of arguments")
	}
	k, v := fields[1], fields[2]
	done, err := core.Set(k, v)
	if err == nil && done {
		*reply = "OK"
		return nil
	} else {
		*reply = ""
		return err
	}
}

func (c CmderService) Get(req string, reply *string) error {
	fields := strings.Fields(req)
	if len(fields) < 2 {
		return errors.New("(error) ERR wrong number of arguments")
	}
	k := fields[1]
	ret, err := core.Get(k)
	if err == nil {
		*reply = ret
		return nil
	} else {
		*reply = ""
		return err
	}
}

func (c CmderService) Keys(req string, reply *string) error {
	fields := strings.Fields(req)
	if len(fields) < 2 {
		return errors.New("(error) ERR wrong number of arguments")
	}
	k := fields[1]
	ret, err := core.Keys(k)
	if err != nil {
		*reply = ""
		return err
	} else {
		*reply = strings.Join(ret, ", ")
		return nil
	}
}

// cmd like "SET name jianxin 60"
// identify wrong command wtih panic
// func splitCmd(cmd string) (k,v string,params []string){
// 	filds := strings.Fields(cmd)
// 	return "","",[]string{""}
// }

const port string = ":8064"

// start the goredis server
func Start() {
	rpc.RegisterName("CmderService", new(CmderService))

	listener, err := net.Listen("tcp", port)
	goredisStartOutput()
	checkErr(err)
	for conn, err := listener.Accept(); err == nil; {
		rpc.ServeConn(conn)
	}
}

// func router(cmdStr string,handler func(cmd string) (string,error)) (string,error) {

// }

func checkErr(err error) {
	if err != nil {
		log.Fatal("Accpter err:", err)
	}
}

func goredisStartOutput() {
	log.Println("goredis is starting")
	log.Printf("goredis server listening on 127.0.0.1%s ....\n", port)
}
