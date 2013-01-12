package main

type ttt interface{}
type tttt []ttt

func main() {
	a := "123456"
	println(a[0:3])
	b := make(tttt, 1)
	b[0] = a
	for k, v := range b {
		v1, ok := v.(string)
		println(k, v1, ok)
	}
}
