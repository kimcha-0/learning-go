package main

import (
    "fmt"
    "time"
)

func main() {

    echo1Start := time.Now()
    echo1()
    echo1End := time.Now()

    echo2Start := time.Now()
    echo2End := time.Now()

    echo3Start := time.Now()
    echo3End := time.Now()

    fmt.Println("Echo test")
    fmt.Println("for loop echo1")
    fmt.Println(echo1End.Sub(echo1Start))
    fmt.Println(echo2End.Sub(echo2Start))
    fmt.Println(echo3End.Sub(echo3Start))

}
