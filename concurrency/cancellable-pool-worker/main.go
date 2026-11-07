package main

import (
	"context"
	"fmt"
	"time"
	"sync"
)

type Job struct {
	ID int
}

type Result struct {
	Job Job
	Result string
}

func worker(ctx context.Context,id int, jobs <-chan Job, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	for{
		select{
		case <- ctx.Done():
			fmt.Println("worker is shutting down and worker id is" , id)
			return 
		case job, ok := <- jobs:
			if !ok{
				fmt.Printf("worker %d : jobs channel closed \n ", id)
				return 
			}

			fmt.Println("worker", id, "processing job ", job.ID)
			
			//simulate work that can be interrupted:
			select {
				case <- time.After(2 * time.Second):
					results <- Result{
						Job: job,
						Result: fmt.Sprintf("worker %d : processed job %d with result", id, job.ID),
					}
				case <- ctx.Done():
					fmt.Printf("worker %d : cancelled while processing job %d \n", id, job.ID)
					return 
			}

		}
	}
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan Job, 20)
	results := make(chan Result, 20)

	var wg sync.WaitGroup

	// Start 3 workers.
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg, results)
	}

	// Produce 20 jobs.
	go func() {
		defer close(jobs)

		for i := 1; i <= 20; i++ {
			select {
			case jobs <- Job{ID: i}:
				fmt.Printf("Main: submitted job %d\n", i)

			case <-ctx.Done():
				fmt.Println("Producer: cancelled")
				return
			}
		}
	}()

	// Cancel after 5 seconds.
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("\nMain: cancelling workers")
		cancel()
	}()

	// Wait for all workers, then close results.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results until results is closed.
	for result := range results {
		fmt.Println("Result:", result.Result)
	}

	fmt.Println("Main: all workers stopped")
}