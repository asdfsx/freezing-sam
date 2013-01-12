package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "test.tsv"
	ostream, err := os.Create(filename)
	defer ostream.Close()
	if err != nil {
		fmt.Printf("err writing file:%v,err:%v\n", filename, err)
	}

	for i := 0; i < 10; i++ {
		stri := strconv.Itoa(i)
		ostream.WriteString(strings.Join([]string{stri, "testtest" + stri + "\n"}, "\t"))
	}
	ostream.Close()
}
