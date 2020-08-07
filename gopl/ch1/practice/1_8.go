package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        if !(strings.HasPrefix("http://", url) && strings.HasPrefix("https://", url)){
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        _, error := io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if error != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, error)
            os.Exit(1)
        }
    }
}
