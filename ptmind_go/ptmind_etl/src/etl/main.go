// ptmind_etl project main.go
package main

import (
	"etlserver"
	"flag"
	"fmt"
	"net/rpc"
	"os"
)

func start(host string, port int) {
	fmt.Println("start the etlserver")
	server := new(etlserver.Etlserver)
	server.NewEtlserver()
	server.Start(port)
}

func stop(host string, port int) {
	fmt.Println("stop the etlserver")
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		fmt.Printf("stop failed, cause:%v\n", err)
		return
	}
	var reply string
	args := &etlserver.RpcArgs{Command: "stop", Value: ""}
	err = client.Call("Rpcserver.RpcProxy", args, &reply)
	if err != nil {
		fmt.Printf("stop failed, cause:%v\n", err)
		return
	}
	fmt.Println("etlserver is stopping")
	client.Close()
}

func status(host string, port int) {
	fmt.Println("get the status of the etlserver")
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		fmt.Printf("stop failed, cause:%v\n", err)
		return
	}
	var reply string
	args := &etlserver.RpcArgs{Command: "status", Value: ""}
	err = client.Call("Rpcserver.RpcProxy", args, &reply)
	if err != nil {
		fmt.Printf("getting status failed, cause:%v\n", err)
		return
	}
	fmt.Printf("etlserver status : %v\n", reply)
	client.Close()
}

func refresh(host string, port int) {
	fmt.Println("refresh the config")
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		fmt.Printf("refresh config failed, cause:%v\n", err)
		return
	}
	var reply string
	args := &etlserver.RpcArgs{Command: "refresh", Value: ""}
	err = client.Call("Rpcserver.RpcProxy", args, &reply)
	if err != nil {
		fmt.Printf("refresh config failed, cause:%v\n", err)
		return
	}
	fmt.Println("etlserver config refreshing")
	client.Close()
}

func redo(host string, port int, redojob string) {
	fmt.Println("refresh the config")
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		fmt.Printf("refresh config failed, cause:%v\n", err)
		return
	}
	var reply string
	args := &etlserver.RpcArgs{Command: "redo", Value: redojob}
	err = client.Call("Rpcserver.RpcProxy", args, &reply)
	if err != nil {
		fmt.Printf("refresh config failed, cause:%v\n", err)
		return
	}
	fmt.Println("etlserver config refreshing")
	client.Close()
}

func main() {
	fmt.Println("Hello World!")
	var action = flag.String("action", "help", "tell me what to do")
	var rpchost = flag.String("host", "127.0.0.1", "rpc host")
	var rpcport = flag.Int("port", 20302, "rpc port")
	var redojob = flag.String("redo", "", "redo job")
	flag.Parse()

	/*
		config, err := config.ReadDefault("/Users/asdfsx/Documents/freezing-sam-git/ptmind_go/ptmind_etl/src/etl/etl.ini")
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		host, _ := config.String("etlrpc", "host")
		port, _ := config.Int("etlrpc", "port")
	*/

	switch *action {
	case "start":
		start(*rpchost, *rpcport)
	case "stop":
		stop(*rpchost, *rpcport)
	case "status":
		status(*rpchost, *rpcport)
	case "refresh":
		refresh(*rpchost, *rpcport)
	case "redo":
		redo(*rpchost, *rpcport, *redojob)
	default:
		fmt.Printf("Usage: %v --action=start|stop|status|refresh|redo --host=127.0.0.1 --port=20302 --redojob=All_1111111111\n", os.Args[0])
	}
}
