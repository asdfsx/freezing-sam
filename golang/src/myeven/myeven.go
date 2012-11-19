package main

import (
    "even"
    "typeExample"
    "helloworld"
    "interfaceExample"
    "listExample"
    "funcExample"
    "structExample"
    "goroutineExample"
    "fmt"
)

func main(){
    i := 5
    fmt.Printf("Is %d even? %v\n", i, even.Even(i))
    helloworld.Helloworld()
    typeExample.TypeExample()
    interfaceExample.InterfaceExample()
    listExample.ListExample()
    funcExample.TestA()
    
    a := []byte{'1','2','3','4',}
    var x int
    for i := 2; i<len(a); {
        x, i = funcExample.NextInt(a, i)
        println(x)
        println(i)
    }

    println("------------------------")
    funcExample.DeferTest1()
    i = funcExample.DeferTest2()
    fmt.Printf("%v\n", i)
    funcExample.MultiArgsTest(1,2,3,4,5)
    funca := func(){
        println("hello world")
    }
    funca()

    funcb := func(s string){
        println(s)
    }
    funcExample.PassFuncTest("hello function",funcb)
    funcExample.FuncMap()
    
    funcc := func(i int){
         println(1/i)
    }
    funcExample.ThrowsPanic(funcc)

    var sb structExample.S
    pa := &sb
    //a := new(structExample.S)
    structExample.F1(pa)
    structExample.F2(pa)
    println(sb.Get())
    println(pa.Get())

    goroutineExample.Gogo()
}
