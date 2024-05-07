package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func fetch() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        // b, err := io.ReadAll(resp.Body)
        b, err := io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %d bytes %s: %v\n", b, url, err)
            os.Exit(1)
        }
        // fmt.Printf("%s", b)
        fmt.Printf("HTTP Status: %s", resp.Status)
    }
}
