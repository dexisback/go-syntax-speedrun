
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID  int
	Output string
	Err    error
}

// Stage 1: Generate jobs.
func generator(ctx context.Context, jobs chan<- Job) {
	defer close(jobs)

	for i := 1; i <= 20; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Generator: cancelled")
			return

		case jobs <- Job{ID: i}:
			fmt.Printf("Generator: submitted job %d\n", i)
		}
	}
}

// Stage 2: Process jobs concurrently.
func worker(
	ctx context.Context,
	id int,
	jobs <-chan Job,
	results chan<- Result,
	cancel context.CancelFunc,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: shutting down\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: no more jobs\n", id)
				return
			}

			// Avoid starting another job after cancellation.
			if ctx.Err() != nil {
				return
			}

			fmt.Printf("Worker %d: processing job %d\n", id, job.ID)

			// Simulate processing.
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d: cancelled during job %d\n", id, job.ID)
				return

			case <-time.After(2 * time.Second):
				// Simulate an error on job 7.
				if job.ID == 7 {
					err := fmt.Errorf("failed to process job %d", job.ID)

					results <- Result{
						JobID: job.ID,
						Err:   err,
					}

					fmt.Printf("Worker %d: encountered an error\n", id)
					cancel()
					return
				}

				result := Result{
					JobID:  job.ID,
					Output: fmt.Sprintf("Job %d processed successfully", job.ID),
				}

				select {
				case <-ctx.Done():
					return
				case results <- result:
				}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)

	// Start the generator.
	go generator(ctx, jobs)

	// Start workers.
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, cancel, &wg)
	}

	// Close results only after every worker has exited.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Stage 3: Collect results.
	for result := range results {
		if result.Err != nil {
			fmt.Printf("Collector: ERROR - %v\n", result.Err)
			continue
		}

		fmt.Printf("Collector: %s\n", result.Output)
	}

	fmt.Println("Pipeline: finished")
}
