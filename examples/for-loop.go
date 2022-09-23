package main

import (
    "fmt"
)

func main() {
    // The only type of Loop in GoLang
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    } 

    fmt.Println("===================")

    // Practically a While loop
    i := 6
    
    for i < 10 {
        fmt.Println(i)
        i++
    }

    fmt.Println("===================")

    // Iterate over an array (fixed length) or a slice (lenght mutable)
    arr := []string{"a", "b", "c"}

    for index, value := range arr {
        fmt.Println("index:", index, "value:", value)
    }

    fmt.Println("===================")

    // Iterate over a map (dictionary)
    m := make(map[string]string)
    m["a"] = "alpha"
    m["b"] = "beta" 

    for key, value := range m {
        fmt.Println("key:", key, "value:", value)
    }
}


