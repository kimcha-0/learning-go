// uses strings.Join to avoid memory consumption from string concat
package main

import (
    "fmt"
    "os"
    "strings"
)

func echo3() {
    fmt.Println(strings.Join(os.Args[:], " "))
}
