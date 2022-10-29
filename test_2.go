package main

import (
	"fmt"
	//"github.com/eiannone/keyboard"
)

var ch1 = make(chan int, 1)
var ch2 = make(chan int, 1)
var ch3 = make(chan int, 1)

func main() {
	ch3 <- 0
	ch1 <- 0
	for {
		<-ch3
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
	fmt.Print("装弹 ->")
	ch2 <- 0
	ch1 <- 0
}
func Miao() {
	fmt.Print("瞄准 ->")
	ch3 <- 0
	ch2 <- 0
}
func Fa() {
	fmt.Print("发射！\n")
	ch1 <- 0
	ch3 <- 0
}
