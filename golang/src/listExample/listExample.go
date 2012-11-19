package listExample

import (
    "fmt"
)

func ListExample(){
    var arr [10]int
    arr[0] = 42
    arr[1] = 13
    fmt.Printf("The first element is %d\n", arr[0])

    a := [2][2]int{[2]int{1,2},[2]int{3,4}}
    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", a[1][1])
    fmt.Printf("=====================\n")

    slice := make([]int, 10)
    fmt.Printf("len(slice)=%v\n", len(slice))
    fmt.Printf("cap(slice)=%v\n", cap(slice))
    fmt.Printf("%v\n", slice[2:5])

    a2 := [...]int{1,2,3,4,5,6,7,8,9}
    s1:= a2[2:4]
    s2:= a2[1:5]
    s3:= a2[:]
    s4:= a2[:4]
    s5:= s2[:]
    fmt.Printf("%v\n", s1)
    fmt.Printf("%v\n", s2)
    fmt.Printf("%v\n", s3)
    fmt.Printf("%v\n", s4)
    fmt.Printf("%v\n", s5)

    s6 := []int{0,0}
    s7 := append(s6, 1)
    s8 := append(s7, 2)
    
    s9 := make([]int,2)
    n := copy(s9, s8[2:4])
    fmt.Printf("%v\n", s6)                                                        
    fmt.Printf("%v\n", s7)                                                        
    fmt.Printf("%v\n", s8)                                                        
    fmt.Printf("%v\n", s9)                                                        
    fmt.Printf("%v\n", n)

    monthdays := map[string]int{
        "Jan":31, "Feb":28, "Mar":31,
        "Apr":30, "May":31, "Jun":30,
        "Jul":31, "Aug":30, "Sep":31,
        "Oct":30, "Nov":31, "Dec":30,
    }

    year := 0
    for _, value := range monthdays{
        year += value
    }
    fmt.Printf("Number of days in a year: %v\n", year)

    monthdays["test"] = 22
    monthdays["Feb"] = 29

    var value int
    var present bool

    value, present = monthdays["test"]
    v, ok := monthdays["test"]

    fmt.Printf("%v, %v\n", value, present)
    fmt.Printf("%v, %v\n", v, ok)
    
    delete(monthdays, "test")
    fmt.Printf("%v\n", monthdays)
    /*for i:=0;i<=100;i++{
        for j:=0;j<=i;j++{
            fmt.Print("A")
        }
        fmt.Print("\n")
    }*/
}
