package main

import (
	"fmt"
	"strconv"
)

// function

func Appender(a int, b string) string {
	return strconv.Itoa(a) + b
}

func KanKun(s string) (st string, l int) {
	st += "verified"
	l = len(s)
	return
}

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

	//slicing like python
	slice := arr2[2:]
	fmt.Println(slice)

	slice = append(slice, "1", "added", "hello")
	fmt.Println(slice)

	// maps

	m := make(map[string]int)
	m["ashin"] = 5
	m["ak"] = 2
	m["none"] = 4

	var val, exist = m["ashin"]
	fmt.Println(val, exist)

	val, exist = m["as"]
	fmt.Println(val, exist)
	fmt.Println(m)

	//string is mutable
	str := "hello go"
	println(str)
	str = "hi there"
	println(str)

	firstcall := Appender(5, "hello")
	fmt.Println(firstcall)

	st, l := KanKun("ashin")
	fmt.Println(st, l)
}
