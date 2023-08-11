package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	contains := slices.Contains(s, 1)
	println(contains) // true

	clone := slices.Clone(s) //Copia by value do array anterior

	slices.Reverse(clone)
	fmt.Println(clone) // 5, 4, 3, 2, 1
	fmt.Println(s)     // 1, 2 , 3, 4, 5

	slices.Sort(clone)
	fmt.Println(clone) //  1, 2 , 3, 4, 5

	m := slices.Max(clone)
	fmt.Println(m) //  5

	m = slices.Min(clone)
	fmt.Println(m) //  1

	d := slices.Delete(clone, 1, 4)
	fmt.Println(d) //  1, 5

	sorted := slices.IsSorted(clone)
	println(sorted) //false

	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found) //Vera: 2 true

}
