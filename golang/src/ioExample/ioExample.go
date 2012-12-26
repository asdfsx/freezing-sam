package ioExample

import(
    "bufio"
    "os"
)

func IoExample(){
    buf := make([]byte, 1024)
    f, _ := os.Open("/etc/passwd")
    defer f.Close()
    r := bufio.NewReader(f)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()
    /*
    for {
        n, _ := f.Read(buf)
        if n == 0 { break }
    */
        /*
        注意下边传入slice后，不需要再对buf进行清空
        原因可以看file_unix.go中的Write函数
        里边有一个b = b[m:n]的操作，这个对逐渐的将buf清空
        传入slice，相当于传入一个指针，函数里对slice的操作会影响到函数外的slice
        如果传入一个数组，只是一个值拷贝，函数理的操作不会对外边产生任何影响
        */
    /*
        os.Stdout.Write(buf[:n])
    }
    */
    println("--------use bufio-------")
    for {
        n, _ := r.Read(buf)
        if n == 0 { break }
        //传入slice，原因同上
        w.Write(buf[0:n])
    }
}
