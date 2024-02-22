package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/atomic-counters
func worker(id int) {
	fmt.Println("worker", id, "started job")
	time.Sleep(time.Second * 5)
	fmt.Println("worker", id, "finished job")
}

func main() {

}
