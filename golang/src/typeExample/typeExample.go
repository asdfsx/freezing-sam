package typeExample

import (
    "fmt"
    "math"
    "strconv"
)

func TypeExample(){
    i := 100000
    fmt.Printf("%v\n" , i*i)
    
    /*
    下边演示了int8无法保存大于128的数据
    var i2 int8
    i2 = 127 + 1
    fmt.Printf("%v\n" , i2)
    */
    
    var i3 int8 = -128
    fmt.Printf("uint8(%v):%v > 127\n", i3, uint8(i3))

    fmt.Println(strconv.FormatUint(1, 2))
    fmt.Println(strconv.FormatUint(math.Float64bits(1.0), 2))

    s := "Go 程"
    r := []rune(s)
    fmt.Printf("%c %c\n", r[0], r[2])
    fmt.Printf("%x\n", r)

    s2 := [4]int{1,2,3,4}
    t := s2[1:3]
    fmt.Println(s2[0], t, s2[:3], t[1:])
    fmt.Println(len(s2), cap(s2))
    fmt.Println(len(t), cap(t))
}
