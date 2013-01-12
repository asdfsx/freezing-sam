package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type test struct {
	id   int
	name string
}

func main() {
	filename := "test.tsv"

	istream, err := os.Open(filename)
	defer istream.Close()
	if err != nil {
		fmt.Printf("file:%v err:%v\n", filename, err)
	}

	r := bufio.NewReader(istream)

	testarray := make([]test, 20)
	i := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			strarr := strings.Split(line, "\t")
			testarray[i].id, _ = strconv.Atoi(strarr[0])
			testarray[i].name = strarr[1]
		}
		i += 1
	}
	println(len(testarray))
	for i = 0; i < len(testarray); i++ {
		fmt.Printf("%v,%v\n", testarray[i].id, testarray[i].name)
	}
}
