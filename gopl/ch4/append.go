package main

import "fmt"

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        z = x[:zlen]
    } else {
        zcap := zlen
        if zcap < 2 * len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[len(x)] = y
    return z
}

func main() {
    var a []int
    fmt.Printf("%d %d %T\n", len(a), cap(a), a)
    a = appendInt(a[0:], 1)
    fmt.Println(a)
    fmt.Printf("%d %d %T\n", len(a), cap(a), a)
}
