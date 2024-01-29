package main

import (
	"fmt"
	"sync"
	"time"
)

// https://gobyexample.com/waitgroups
func worker(id int) {
	fmt.Println("worker", id, "started job")
	time.Sleep(time.Second * 5)
	fmt.Println("worker", id, "finished job")
}

func main() {

	const numJobs = 5

	// define
	var wg sync.WaitGroup

	// loop to create subroutines
	for w := 1; w <= numJobs; w++ {
		wg.Add(1)

		// localize value to pass in go subroutine
		i := w

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// wait for all tasks
	wg.Wait()

	fmt.Println("Hello world!")

}
