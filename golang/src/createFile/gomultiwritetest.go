package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createContent(c chan int){
	rand.Seed( time.Now().UTC().UnixNano())
        c <- rand.Intn(10)
}

func output(c chan int){
        var num int
	for{
		select{
			case num = <- c:
                        	fmt.Printf("%v\n", num)
		}
	}
}

func main() {
        var chanarray []chan int
	for i := 0; i< 4; i=i+1 {
                tmp := make(chan int)
		chanarray = append(chanarray, tmp)
	}
	c := make(chan int)
        go output(c)
        go createContent(c)
        go createContent(c)
	go createContent(c)
	go createContent(c)
	time.Sleep(1000000)    
}
