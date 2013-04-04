package main

import "fmt"

func main() {
	aa := make([][]string, 0)
	aa = append(aa, []string{"1", "b"})
	aa = append(aa, []string{"1", "b"})
	aa = append(aa, []string{"1", "b"})
	aa = append(aa, []string{"1", "b"})
	aa = append(aa, []string{"1", "b"})
	aa = append(aa, []string{"1", "b"})
	aa = append(aa, []string{"1", "b"})
	fmt.Printf("%v\n", aa)
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	if true {
		println(monthdays["Jan"])
	}
        v, ok := monthdays["1"]
        println(v)
        println(ok)
}
