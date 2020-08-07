package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var filenames []string

    files := os.Args[1:]
    if len(files) == 0 {
        countLines("", os.Stdin)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            if countLines(arg, f) {
                filenames = append(filenames, arg)
            }
            defer f.Close()
        }
    }

    fmt.Println(filenames, len(filenames))
    for _, filename := range filenames {
        fmt.Println("filename: ", filename)
    }
}

func countLines(filename string, f *os.File) bool {
    input := bufio.NewScanner(f)
    lines := make(map[string]int)
    for input.Scan() {
        lines[input.Text()] += 1
    }

    var r bool = false
    for _, n := range lines {
        if n > 1 {
            r = true
            break
        }
    }
    return r
}
