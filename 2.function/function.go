package main

import "fmt"

func main() {
	fmt.Println("1+2=")
	fmt.Println(add(1, 2))
}

func add(a, b int) int {
	return a + b
}
