package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	mapa := make(map[int]int)
	var key1 = 0
	for ; n > 0; n-- {
		fmt.Scan(&key1)
		if mapa[key1] != 0 {
			mapa[key1] += 1
		} else {
			mapa[key1] = 1
		}
	}
	for a, b := range mapa {
		fmt.Printf("%d %d\n", a, b)
	}
}
