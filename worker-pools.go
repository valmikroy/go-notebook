package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/worker-pools
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second * 5)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create workers - creating only 3 workers to process 5 jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send workers task
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect result of task
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	fmt.Println("Hello world!")

}
