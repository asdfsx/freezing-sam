package main

import (
"os"
"fmt"
"bufio"
"path/filepath"
)

func readfile(filename string, c chan int){
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
            fmt.Printf("%v", line)
        }
    }
    c <- 1
}

func getfilelist(path string)(filelist []string, err error){
    filelist, err = filepath.Glob(path)
    return
}

func main(){
    path := "/Users/asdfsx/Documents/freezing-sam-git/golang/src/createFile/*"
    filelist, err := getfilelist(path)
    if err != nil {
        fmt.Printf("filepath:%v, error:%v\n", path, err)
    }

    filenum := len(filelist)

    c := make(chan int)
    for _, file := range(filelist) {
        println(file)
        go readfile(file, c)
    }
    i := 0
    L: for {
        select {
        case result := <- c:
            i++
            println(result)
            if i > filenum-1 {
                break L
            }
        }
    }
}
