// Echo2 prints its command-line args
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""

    // can also init variables
    // var s string
    // var s = ""
    // var s string = ""

    // range produces a pair: index, value at index
    for i, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
        fmt.Print(i)
        fmt.Print(sep + s + "\n") 
    }
    fmt.Println(s)
}
