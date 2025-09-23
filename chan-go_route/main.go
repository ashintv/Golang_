package main

import (
	"fmt"
	"time"
)

func routune_one(c chan string) {
	for true {
		c <- "routune_one"
		time.Sleep(1 * time.Second)
	}
}

func routune_two(c chan string) {
	for true {
		c <- "routune_two"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Starting =====================")
	c := make(chan string)
	for i:=0 ; i<20 ;i++ {
		go routune_one(c)
		go routune_two(c)
		val := <-c
		fmt.Println(val)

	}
}
