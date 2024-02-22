package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const maxCapacity int = 64 * 1024

func fileReadline(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		// line by line iteration
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileReadline("waitgroup2.go")
}
