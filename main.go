package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	numJobs, err := strconv.Atoi(os.Getenv("NUM_JOBS"))
	if err != nil {
		fmt.Println("Error parsing NUM_JOBS")
	}

	numWorkers, err := strconv.Atoi(os.Getenv("NUM_WORKERS"))
	if err != nil {
		fmt.Println("Error parsing NUM_WORKERS")
	}

	jobs := make(chan int, numJobs)
	results := make(chan int, numWorkers)

	for w := 1; w <= int(numWorkers); w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
