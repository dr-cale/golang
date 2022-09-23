package main

import (
    "fmt"
)

func main() {
    var x int = 6      // falls back to zero if none provided
    var y int = 7
    var sum int = x + y      

    // GO can also infer the type, no need for explicit syntax
    z := 2
    sum = sum + z

    fmt.Println(sum)
}

