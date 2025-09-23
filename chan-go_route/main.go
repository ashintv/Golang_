package main

import (
	"fmt"
	"sync"
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

func routune_three(c chan string, i int, wc *sync.WaitGroup) {
	defer wc.Done()
	time.Sleep(1 * time.Second)
	c <- fmt.Sprintf("routine %d done", i)
	
}

func main() {
	fmt.Println("Starting =====================")

	// unbefferd channel
	c := make(chan string)
	go routune_one(c)
	go routune_two(c)
	for i := 0; i < 5; i++ {
		// this will block until a new value comes
		// blocks until on of routiens send back
		val := <-c
		fmt.Println(val)
	}
	fmt.Println("Starting ===================== buffered")
	c2 := make(chan string, 5)

	go routune_one(c2)
	go routune_two(c2)
	for i := 0; i < 10; i++ {
		// wait until 5 values come
		val := <-c2
		fmt.Println(val)
	}
	fmt.Println("Starting ===================== wt goups")
	//wait grops
	// this something like semafor count it will icreased when a new threead starts and decreases when thread completes like semaphore count

	var wc sync.WaitGroup
	//creating 5 buffer channel
	c3 := make(chan string, 5)
	for i := 0; i < 5; i++ {
		wc.Add(1)
		go routune_three(c3, i, &wc)
	}
	wc.Wait()
	close(c3)
	for val := range c3 {
		fmt.Println(val)
	}

}
