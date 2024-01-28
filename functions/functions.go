package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x, y := swap("hello", "world")
	fmt.Println(x, y)
}

// :other way
// func add(x, y int) int {
// 	return x + y
// }

// :base
// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }
