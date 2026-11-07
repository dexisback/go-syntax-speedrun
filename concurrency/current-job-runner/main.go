//go worker pool
//main() creates jobs and collects results
//jobs channel -> ex: 20 jobs waiting to be processed
//then spawn 5 workers (5 goroutines) who actually do the work
//results channel -> collects the results from all these workers 
//main() reads and prints results

package main 



import (
	"fmt"
	"sync"
	"time"
)


type Job struct {
	ID int 
}


type Result struct {
	JobID int 
	Output string 
}

// <- means receive only
//chan <- means send only


//worker needs an id, a job channel(3 jobs we'll be running), and results channel idk why, and wg for controlling
func worker(
	id int , 
	jobs <-chan Job, 	//can only receive jobs
	results chan <- Result,   //can only send results
	wg *sync.WaitGroup) {
		
		defer wg.Done()

		for job := range jobs {
			fmt.Printf("worker %d is processing job %d\n", id, job.ID)
			time.Sleep(time.Second * 2)

			results <- Result {
				JobID: job.ID,
				Output: fmt.Sprintf("Job %d completed", job.ID),
			}
		}

	}


func main(){
	jobs := make(chan Job , 5)
	results := make(chan Result, 20)

	var wg sync.WaitGroup

	//Todo1: launch 3 workers
	//Todo2: send 20 jobs into jobs channel
	//Todo3: close jobs after sending all jobs 
	//Todo4: wait for all workers to finish
	//Todo5: close results channel
	//Todo6: read and print all results 
	wg.Add(3)
	//launch 3 workers
	for i := 0 ; i < 3 ; i++ {
		go worker(i, jobs, results, &wg)
	}
	//20 jobs:
	for i := 0 ; i<20 ; i++ {
		jobs <- Job{
			ID: i,
		}
	}
	close(jobs)
	wg.Wait() //wait until the counter is 0 (so all workers have finished)
	
	close(results)

	for result := range results {
		fmt.Println("Result is :", result.Output)
	}

}	