package main

import (
	"fmt"
	//"github.com/eiannone/keyboard"
)

var ch1 = make(chan int, 1)
var ch2 = make(chan int, 1)
var ch3 = make(chan int, 1)

//var quit = make(chan int, 1)
//var sum = 0

func main() {
	ch3 <- 0
	ch1 <- 0
	//var char rune
	//k := &char
	//go ck(k)
	for {
		<-ch3
		//if char == 'q' {
		//	quit <- 0
		//	sum += 1
		//}
		//if sum == 3 {
		//	break
		//}
		select {
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		case <-ch1:
			go Zhuang()
		default:

		}
		<-ch1
		//if char == 'q' {
		//	quit <- 0
		//	sum += 1
		//}
		//if sum == 3 {
		//	break
		//}
		select {
		case <-ch2:
			go Miao()
		case <-ch2:
			go Miao()
		case <-ch2:
			go Miao()
		case <-ch2:
			go Miao()
		case <-ch2:
			go Miao()
		default:

		}
		<-ch2
		//if char == 'q' {
		//	quit <- 0
		//	sum += 1
		//
		//}
		//if sum == 3 {
		//	break
		//}
		select {
		case <-ch3:
			go Fa()
		case <-ch3:
			go Fa()
		case <-ch3:
			go Fa()
		default:

		}
	}
}
func Zhuang() {
	//select {
	//case <-quit:
	//	return
	//
	//}
	fmt.Print("装弹 ->")
	ch1 <- 0
	ch2 <- 0
}
func Miao() {
	//select {
	//case <-quit:
	//	return
	//
	//}
	fmt.Print("瞄准 ->")
	ch2 <- 0
	ch3 <- 0
}
func Fa() {
	//select {
	//case <-quit:
	//	return
	//
	//}
	fmt.Print("发射！\n")
	ch3 <- 0
	ch1 <- 0
}

//func ck(a *rune) {
//	char, _, err := keyboard.GetSingleKey()
//	if err != nil {
//		panic(err)
//	}
//	*a = char
//}
