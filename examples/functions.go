package main

import (
    "fmt"
    "strconv"
    "errors"
    "math"
)

func main() {
    result := sum(2, 3)
    string_result := strconv.Itoa(result)
    fmt.Println("The sum of 2 and 3 is", string_result + ".")

    number, err := sqrt(16)
    string_number := strconv.FormatFloat(number, 'f', -1, 32)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("The square root of number 16 is " + string_number + ".")
    }
}

func sum(x int, y int) int {
    return x + y
}

func sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("Square root is undefined for negative numbers.")
    }

    return math.Sqrt(x), nil
}


