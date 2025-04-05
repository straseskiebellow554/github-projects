package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    _, err := fmt.Scan(&num)
    if err != nil {
        panic(err)
    }
    if num%2 == 0 {
        fmt.Println("The number is even.")
    } else {
        fmt.Println("The number is odd.")
    }
}
