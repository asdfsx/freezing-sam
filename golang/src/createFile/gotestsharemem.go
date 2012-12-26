package main

import (
"os"
"fmt"
"bufio"
"strings"
"time"
)

func readfile(filename string, dict map[string]string){
    println(dict)
    istream, err := os.Open(filename)
    defer istream.Close()
    if err != nil{
        println(err)
        return
    }
    r := bufio.NewReader(istream)

    for {
        line, err := r.ReadString('\n')
        if err != nil {
            println(err)
            break
        } else {
            strarr := strings.Split(line, "\t")
            dict[strarr[0]] = strarr[1]
        }
    }
}

func readdict(dict *map[string]string){
    for k, v := range *dict {
        fmt.Printf("%v:%v", k, v)
    }
}

func main(){
    path := "/Users/asdfsx/Documents/freezing-sam-git/golang/src/createFile/test.tsv"

    dict := make(map[string]string)
    println(dict)

    readfile(path, dict)
    go readdict(&dict)
    /*for k, v := range dict {
        fmt.Printf("%v:%v\n", k, v)
    }*/
    time.Sleep(100)
}
