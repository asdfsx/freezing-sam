package main

import (
"fmt"
)

func A () [][]string {
    t := make([][]string, 0)
    a := []string{"1","2","3","4","5"}
    t = append(t, a)
    return t
}

func main(){
    b := A()
    fmt.Printf("%v\n", b)
}
