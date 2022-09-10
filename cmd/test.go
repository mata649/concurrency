package main

import (
	"fmt"
	"time"
)

func testFunc(id int, c chan int, r chan int) {
	fmt.Println("Starting test func with id", id)
	time.Sleep(time.Second * 3)
	fmt.Println("Ending test func with id", id)
	c <- id * 2

}

func printResults() {

}
func main() {
	c := make(chan int, 1)
	results := make(chan int, 1)

	for i := 1; i <= 10; i++ {

		go testFunc(i, c, results)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Before c print")
		fmt.Println(<-c)
		fmt.Println("After c print")

	}

}
