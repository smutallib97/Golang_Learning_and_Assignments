package main

import (
	"fmt"
	"time"
)

func worker(id int, job <-chan string, results chan<- string) {
	for task := range job {
		fmt.Printf("Worker %d started task %s\n", id, task)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished task %s\n", id, task)
		results <- task
	}
}

func main() {
	jobs := make(chan string, 100)
	results := make(chan string, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	jobs <- "Task 1"
	jobs <- "Task 2"
	jobs <- "Task 3"

	close(jobs)

	for i := 1; i <= 3; i++ {
		fmt.Println("Task", <-results, "is done.")
	}
}
