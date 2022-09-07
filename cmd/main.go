package main

import "fmt"

// func main() {
// 	c := make(chan int, 3)
// 	c <- 1
// 	c <- 2
// 	c <- 3
// 	c <- 4
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }

func main() {
	c := make(chan int)
	go escuchar(c)
	publicar(c)
}

func escuchar(c <-chan int) {
	for {

		fmt.Println(<-c)

	}
}

func publicar(c chan<- int) {
	c <- 2
	c <- 4
	c <- 5

	fmt.Println("Funcion publicar")
}
