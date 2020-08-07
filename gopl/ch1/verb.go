package main

import (
    "fmt"
    "math"
)

func main() {
    var i int = 10
    fmt.Printf("十进制 %d %T\n", i, i)
    fmt.Printf("十六进制 %x\n", i)
    fmt.Printf("八进制 %o\n", i)
    fmt.Printf("二进制 %b\n", i)

    var f float64 = math.Pi
    fmt.Printf("%f\n", f)
    fmt.Printf("%g\n", f)
    fmt.Printf("%e\n", f)

    var tb bool = true
    fmt.Printf("%t\n", tb)

    var fb bool = false
    fmt.Printf("%t\n", fb)

    var s string = "Hello, world"
    fmt.Printf("%s\n", s)
    fmt.Printf("%q\n", s)

    type video struct {
        ID int
    }
    var v video
    fmt.Printf("%v %T\n", v, v)

    fmt.Printf("%%\n")
}
