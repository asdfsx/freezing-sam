package etlserver

import (
	"config"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

type Etlserver struct {
	Etlconf *config.Config
	rpcchan chan string
	Rpcserv *Rpcserver
}

func (self *Etlserver) NewEtlserver() {
	self.rpcchan = make(chan string)
	self.Rpcserv = &Rpcserver{Rpcchan: self.rpcchan}
}

func (self *Etlserver) Start(rpcport int) {
	fmt.Printf("eltserver start at port: %v\n", rpcport)
	self.startRpcServ(rpcport)
	self.run()
}

func (self *Etlserver) startRpcServ(rpcport int) {
	rpc.Register(self.Rpcserv)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", fmt.Sprintf(":%v", rpcport))
	if e != nil {
		panic(e)
	}
	fmt.Printf("rpc server start\n")
	go http.Serve(l, nil)
}

func (self *Etlserver) run() {
	var content string
	for {
		select {
		case content = <-self.rpcchan:
			switch {
			case content == "stop":
				goto QUIT
			case content == "status":
				self.rpcchan <- "12345667890"
			case content == "refresh":
				fmt.Println("refreshing")
			case strings.HasPrefix(content, "redo:"):
				fmt.Printf("%v\n", content)
			default:
				fmt.Printf("cannot find proper command: %v\n", content)
			}
		}
	}
QUIT:
	fmt.Println("finish running")
}
