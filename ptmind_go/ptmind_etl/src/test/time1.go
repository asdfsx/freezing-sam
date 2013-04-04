package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()
	fmt.Println(t)

	fmt.Println(time.Unix(t, 0).String())
	fmt.Println(time.Now().String())

	//带纳秒的时间戳
	t = time.Now().UnixNano()
	fmt.Println(t)
	fmt.Println(time.Unix(0, t).String())
	fmt.Println("------------------")

	//基本格式化的时间表示
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().Format("20060102"))

	t1, _ := time.Parse("200601021504", "201302142006")
	fmt.Println(t1.String())
	fmt.Println(t1.Unix())
}
