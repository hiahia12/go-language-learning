package main

import "fmt"

func main() {
	var name string
	var ufo string
	var p int = 1
	var q int = 1
	fmt.Scanln(&ufo)
	fmt.Scanln(&name)
	var n int = len(name) - 1
	var m int = len(ufo) - 1
	for ; n >= 0; n-- {
		a := name[n]
		t := int(a) - 64
		p = p * t
	}
	for ; m >= 0; m-- {
		a := ufo[m]
		t := int(a) - 64
		q = q * t
	}
	if q%47 == p%47 {
		fmt.Println("GO")
	} else {
		fmt.Println("STAY")
	}

}
