package etlserver

import (
	"fmt"
)

type Rpcserver struct {
	Rpcchan chan string
}

type RpcArgs struct {
	Command string
	Value   string
}

func (self *Rpcserver) RpcProxy(args *RpcArgs, reply *string) error {
	switch args.Command {
	case "stop":
		self.Rpcchan <- "stop"
		*reply = "stopping"
		return nil
	case "status":
		self.Rpcchan <- "status"
		status := <-self.Rpcchan
		*reply = status
		return nil
	case "refresh":
		self.Rpcchan <- "refresh"
		*reply = "refreshing"
		return nil
	case "redo":
		self.Rpcchan <- fmt.Sprintf("%v:%v", "redo", args.Value)
		*reply = "redoing"
		return nil
	default:
		*reply = "unknow command:" + args.Command
		return nil
	}
	return nil
}
