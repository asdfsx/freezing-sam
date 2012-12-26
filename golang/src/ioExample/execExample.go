package ioExample

import (
    "os/exec"
    "fmt"
)

func ExecExample(){
    cmd := exec.Command("/bin/ls", "-l")
    err := cmd.Run()
    buf, err2 := cmd.Output()
    fmt.Printf("%v\n%v\n%v\n", err, err2, buf)
}
