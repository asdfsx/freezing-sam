package main

import (
	"fmt"
)

func main() {
	ss := make([]int, 2)
	ss[0] = 1
	ss[1] = 2
	fmt.Printf("%v\n", ss)
        a := len(ss)
        b := ss[0:a+1]
        b[2] = 3
}
