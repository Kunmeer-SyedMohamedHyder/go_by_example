package main

import (
	"fmt"
)

func work(id int, jobs <-chan int, results chan<- int) {

	fmt.Println("Worker", id, "started")
	for job := range jobs {
		fmt.Println("Worker", id, "doing job num", job)
		results <- job * 2
	}
}

func create_jobs(jobs chan<- int, num_jobs int) {
	for job := 0; job < num_jobs; job++ {
		jobs <- job
	}
}

func main() {
	num_workers := 3
	num_jobs := 10

	jobs := make(chan int, num_jobs)
	results := make(chan int, num_jobs)
	defer close(results)

	for worker := 0; worker < num_workers; worker++ {
		go work(worker, jobs, results)
	}

	create_jobs(jobs, num_jobs)
	close(jobs)

	for r := 0; r < num_jobs; r++ {
		fmt.Println("result", <-results)
	}

}
