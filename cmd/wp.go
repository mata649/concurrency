package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d \n", id, job, fib)
		results <- fib
	}
}
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{2, 3, 4, 5, 23, 14, 40}
	nWorkers := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go worker(i, jobs, results)
	}

	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 3
	jobs <- 3
	jobs <- 3
	jobs <- 3

	// for _, value := range tasks {
	// 	jobs <- value
	// }
	// close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}
	close(results)
}
