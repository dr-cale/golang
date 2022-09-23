package main

import (
    "fmt"
)

type person struct {
    name string
    age int
}

func main() {
    p := person{name: "Jake", age: 33}

    fmt.Println(p)
}


