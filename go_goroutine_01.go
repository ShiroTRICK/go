package main

import (
	"fmt"
	"time"
)

func x() {
	for i := 1; i < 10; i++ {
		fmt.Println("x")
	}
}

func y() {
	for i := 1; i < 10; i++ {
		fmt.Println("y")
	}
}

func z() {
	fmt.Println("z")
}

func k() {
	fmt.Println("k")
}

func main() {
	defer z()
	defer k()
	go x()
	go y()
	time.Sleep(time.Second * 2)
	fmt.Println("end main()")
}
