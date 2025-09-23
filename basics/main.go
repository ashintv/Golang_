// Go Course Template - Comprehensive Golang Fundamentals
// This file demonstrates core Go language concepts including:
// - Basic syntax and control structures
// - Functions and methods
// - Structs and interfaces
// - Goroutines and concurrency
// - Data structures (arrays, slices, maps)

package main

import (
	"fmt"     // Package for formatted I/O operations
	"strconv" // Package for string conversions
	"time"    // Package for time-related operations
)

// FUNCTIONS SECTION
// Functions are first-class citizens in Go and can be passed around as values

// Expensive demonstrates a long-running operation that simulates heavy computation
// This function is designed to be used with goroutines to show concurrent execution
func Expensive(s string) {
	// Loop through 100 iterations, printing the string and counter
	for i := range 100 {
		fmt.Println(s, i)
		time.Sleep(1 * time.Second) // Simulate work by sleeping for 1 second
	}
}

// Appender demonstrates basic function parameters and string operations
// Takes an integer and string, concatenates them and returns the result
func Appender(a int, b string) string {
	return strconv.Itoa(a) + b // Convert int to string and concatenate
}

// KanKun demonstrates named return values - a Go feature that allows
// you to declare return variable names in the function signature
func KanKun(s string) (st string, l int) {
	st += "verified" // Named return variable automatically initialized
	l = len(s)       // Calculate length of input string
	return           // Naked return - returns named variables
}

// HIGHER-ORDER FUNCTIONS
// CallB demonstrates functions as first-class citizens in Go
// It accepts another function as a parameter and uses it
func CallB(f1 func(a int, b string) string, b string) string {
	srt := f1(5, "hello") + b // Call the passed function with arguments
	return srt
}

// STRUCTS AND METHODS SECTION
// rect1 demonstrates a custom struct type for representing a rectangle
type rect1 struct {
	len int // length of rectangle
	wid int // width of rectangle
}

// Findarea demonstrates a function that operates on a struct
// This is a regular function (not a method) that takes a struct as parameter
func Findarea(rect rect1) int {
	return rect.len * rect.wid
}

// AttachedArea demonstrates a method attached to the rect1 struct
// Methods in Go are functions with a special receiver argument
func (r rect1) AttachedArea() int {
	return (r.len + r.wid) * 2 // Calculate perimeter
}

// Updater demonstrates a method with a pointer receiver
// Pointer receivers allow the method to modify the struct it's called on
func (r *rect1) Updater(l int, w int) {
	r.len = l // Modify the original struct (not a copy)
	r.wid = w
}

// MAIN FUNCTION - Entry point of the program
func main() {
	// CONTROL STRUCTURES AND LOOPS SECTION

	// Traditional for loop with initialization, condition, and increment
	for i := 1; i < 10; i++ {
		println(i)
	}

	// Range-based for loop - modern Go syntax for iterating over ranges
	for i := range 10 {
		println("loop2", i)
	}

	// BOOLEAN LOGIC AND CONDITIONAL STATEMENTS
	Boo := true // Variable declaration with type inference

	// Basic if statement
	if Boo {
		println("Heelo", Boo)
		Boo = false // Modify variable inside if block
	}

	// If statement with negation operator
	if !Boo {
		println("Not Boo", Boo)
		Boo = true
	}

	// WHILE-LIKE LOOP USING FOR
	// Go doesn't have a while loop, but for can be used as while
	c := 0
	for Boo { // Continue while Boo is true
		println("Count", c)
		c++
		if c == 10 {
			println("Done")
			Boo = false // Break condition
		}
	}

	// SWITCH STATEMENTS
	// Demonstrates switch cases with fallthrough behavior
	for i := range 4 {
		println("Logging for day no", i)
		switch i {
		case 1:
			println("Day 1")
		case 2:
			println("Day 2")
			fallthrough // Continues execution into the next case
		case 3:
			println("Day 3") // This executes for both case 2 and case 3
		}
	}

	// ARRAYS AND SLICES SECTION
	// Array: Fixed-size sequence of elements
	var arr [5]int = [5]int{1, 2, 3, 4, 5} // Array with explicit size
	var arr2 []string = []string{}         // Slice: Dynamic array
	fmt.Println(arr)

	// Range loop over array - demonstrates iteration and slice manipulation
	for _, value := range arr {
		fmt.Println(value)
		// Convert integer to string and append to slice
		arr2 = append(arr2, strconv.Itoa(value)+"Number")
	}
	fmt.Println(arr2)

	// SLICE OPERATIONS
	// Slicing demonstrates Python-like array slicing syntax
	slice := arr2[2:] // Take elements from index 2 to end
	fmt.Println(slice)

	// Dynamic slice expansion using append function
	slice = append(slice, "1", "added", "hello")
	fmt.Println(slice)

	// MAPS (HASH TABLES/DICTIONARIES)
	// Maps are Go's built-in key-value data structure
	m := make(map[string]int) // Create a map with string keys and int values
	m["ashin"] = 5            // Add key-value pairs
	m["ak"] = 2
	m["none"] = 4

	// Map lookup with existence check - Go's comma ok idiom
	var val, exist = m["ashin"] // val gets the value, exist is true if key exists
	fmt.Println(val, exist)     // Should print: 5 true

	val, exist = m["as"]    // Key doesn't exist
	fmt.Println(val, exist) // Should print: 0 false (zero value for int)
	fmt.Println(m)

	// STRING MUTABILITY
	// Strings in Go are immutable, but variables can be reassigned
	str := "hello go"
	println(str)
	str = "hi there" // This creates a new string, doesn't modify the original
	println(str)

	// FUNCTION CALLS AND DEMONSTRATIONS

	// Call simple function with parameters
	firstcall := Appender(5, "hello")
	fmt.Println(firstcall)

	// Call function with named return values
	st, l := KanKun("ashin") // Multiple assignment from function return
	fmt.Println(st, l)

	// Call higher-order function (function that takes another function as parameter)
	res := CallB(Appender, "Ashin") // Pass Appender function as argument
	fmt.Println(res, "as callB result")

	// STRUCT METHODS AND OPERATIONS

	// Create struct instance using field names (recommended approach)
	rect_1 := rect1{len: 5, wid: 5}

	// Call function that takes struct as parameter
	area := Findarea(rect_1)
	fmt.Println(area)

	// Call method with pointer receiver to modify the struct
	rect_1.Updater(10, 10)

	// Call method on the struct - demonstrates method syntax
	fmt.Println("  ", rect_1.AttachedArea())

	// GOROUTINES - CONCURRENT PROGRAMMING
	// Goroutines are lightweight threads managed by Go runtime
	go Expensive("Ashin calling") // Launch first goroutine
	go Expensive("Ak calling")    // Launch second goroutine

	// Sleep the main thread to prevent program from exiting
	// Without this, main would exit before goroutines complete
	// In production code, use sync.WaitGroup or channels for coordination
	time.Sleep(10 * time.Second)
}
