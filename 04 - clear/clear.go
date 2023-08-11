package main

import "fmt"

func main() {
	// Map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	clear(m)
	fmt.Println(m) // Outputs: map[]

	// Slice
	s := []int{1, 2, 3, 4, 5}
	clear(s)
	fmt.Println(s) // Outputs: [0 0 0 0 0]
}
