package main

import(
	"fmt"
	"flag"
	"os"
)

func start(){
	fmt.Println("start")
}

func stop(){
	fmt.Println("stop")
}

func status(){
	fmt.Println("status")
}

func refresh(){
	fmt.Println("refresh")
}

func main(){
	action := flag.String("action", "help", "tell me what to do")
	flag.Parse()

	switch *action {
		case "start":
			start()
		case "stop":
			stop()
		case "status":
			status()
		case "refresh":
			refresh()
		default:
			fmt.Printf("usage: %s --action=start|stop|status|refresh\n", os.Args[0])
	}
}
