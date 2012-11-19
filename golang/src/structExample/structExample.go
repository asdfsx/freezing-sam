package structExample

import "fmt"

type S struct{
    i int
}
func (p *S) Get()(int){ return p.i }
func (p *S) Put(i int){ p.i = i }

type I interface{
    Get()(int)
    Put(int)
}

type R struct{
    i int
}
func (p *R)Get()(int){ return p.i }
func (p *R)Put(i int){ p.i = i }

func F1(p I){
    fmt.Printf("%v\n", p.Get())
    p.Put(1)
}

func F2(p I){
    switch p.(type){
        case *S:
            println("type *S")
        case *R:
            println("type *R")
        //case S:
        //    println("type S")
        //case R:
        //    println("type R")
        default:
            println("nothing")
    }
}
