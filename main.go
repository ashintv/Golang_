package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for i := 1; i < 10; i++ {
	// 	println(i)
	// }

	// for i := range 10 {
	// 	println("loop2", i)
	// }

	// Boo := true

	// if Boo {
	// 	println("Heelo", Boo)
	// 	Boo = false
	// }

	// if !Boo {
	// 	println("Not Boo", Boo)
	// 	Boo = true
	// }

	// c := 0
	// for Boo {
	// 	println("Count", c)
	// 	c++
	// 	if c == 10 {
	// 		println("Done")
	// 		Boo = false
	// 	}
	// }

	// for i := range 4 {
	// 	println("Logging for day no" , i)
	// 	switch i {
	// 	case 1:
	// 		println("Day 1")
	// 	case 2:
	// 		println("Day 2")
	// 		fallthrough
	// 	case 3:
	// 		println("Day 3")

	// 	}
	// }

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 []string = []string{}
	fmt.Println(arr)

	for _, value := range arr {
		fmt.Println(value)
		arr2 = append(arr2, strconv.Itoa(value)+"Number")
	}
	fmt.Println(arr2)
}
