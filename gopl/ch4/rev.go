package main

import "fmt"

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    a := [...]int{0, 1, 2, 3, 4, 5}
    reverse(a[:])
    fmt.Println(a)

    b := [...]int{0, 1, 2, 3, 4, 5}
    reverse(b[:2])
    fmt.Println(b)
    reverse(b[2:])
    fmt.Println(b)
    reverse(b[:])
    fmt.Println(b)
}
