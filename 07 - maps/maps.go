package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	clone := maps.Clone(m)

	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 != 0 // delete odd values
	})
	fmt.Println(m) //map[four:4 two:2]

	fmt.Println(clone) //map[four:4 one:1 three:3 two:2]

	//maps.Keys()
	//maps.Values()

	//	https://pkg.go.dev/maps@go1.21rc2
}
