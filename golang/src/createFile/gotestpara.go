package main

import "fmt"

func test(a int){
    println(a)
    println(&a)
}

func test2(a string){
    println(a)
    println(&a)
}

func test3(a [3]int){
    fmt.Printf("%v\n",a)
    a[0] = 10
    fmt.Printf("%v\n",a)
}

func test4(a []int){
    fmt.Printf("%v\n",a)
    a[0] = 10
    fmt.Printf("%v\n",a)
}

func test5(a map[string]string){
    fmt.Printf("%v\n",a)
    a["a"] = "asdf"
    fmt.Printf("%v\n",a)
}

func main(){
/*
    a := 1
    println(a)
    println(&a)
    println("--------------")
    test(a)

    b := "test"
    println(b)
    println(&b)
    println("-------------")
    test2(b)

    c := [3]int{1,2,3}
    println(c)
    fmt.Printf("%v\n", c)
    test3(c)
    fmt.Printf("%v\n",c)

    d := []int{1,2,3}
    fmt.Printf("%v\n",d)
    println("-------------")
    test4(d)
    fmt.Printf("%v\n",d)
*/
    println("=============")
    e := map[string]string{"a":"a","b":"b","c":"c"}
    fmt.Printf("%v\n",e)
    println("-------------")
    test5(e)
    fmt.Printf("%v\n",e)
}
