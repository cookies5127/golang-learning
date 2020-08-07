package main

import "fmt"

func nonempty2(strings []string) []string {
    out := strings[:0]
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}

func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i += 1
        }
    }
    return strings[:i]
}

func main() {
    s := []string{"one", "", "three"}
    n := nonempty(s)
    fmt.Println(n, s)
}
