package main

import (
    "fmt"
)

func main() {
    var a [5]int        // An array of 5 integers
    a[2] = 7

    b := [5]int{5, 4, 3, 2, 1}  // Immutable length
    c := []int{6, 4, 1, 3, 8}   // Mutable length
    
    d := append(c, a[2])

    fmt.Println(d)      // Prints all zeroes if unassigned
    fmt.Println(b)
}

