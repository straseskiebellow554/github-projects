package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	for i, name := range names {
		fmt.Println(i, name)
	}
}
