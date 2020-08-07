package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    for i, v := range os.Args[1:] {
        fmt.Println("index: ", i, " value: ", v)
    }
}
