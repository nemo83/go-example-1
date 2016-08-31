package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(swap(1, 2))
}

