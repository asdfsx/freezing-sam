package main

import "path/filepath"

func main(){
    list, _ := filepath.Glob("./*")
    for _, v := range list {
        println(v)
    }
}
