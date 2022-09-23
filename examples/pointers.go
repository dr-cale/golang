package main

import (
    "fmt"
)

func main() {
    i := 7
    fmt.Println(i)    // prints the value of the variable
    inc(i)            // ---> this has no effect
    fmt.Println(i)    // prints the value of the variable
    inc(i)            // ---> this has no effect
    incr(&i)          // ---> this works
    fmt.Println(i)    // prints the value of the variable
    fmt.Println(&i)   // prints the memory address of the variable (pointer to "i")
}

// The "i" copied by value, so the Increment function gets a copy of "i"
// Nothing is returned and the new value is discarded
func inc(x int) {
    x++
}

func incr(x *int) {
    *x++
}

