Goroutines and WaitGroups
Unbuffered vs buffered channels
Producer-consumer pattern
select and default
Channel closing and range
Blocking behaviour

job processing system. 

Imagine you have 20 jobs that need processing. Instead of processing them sequentially, we'll launch 3 workers that consume jobs concurrently.


                 PRODUCER
                    |
                    | sends jobs
                    v
               jobs channel
                    |
          +---------+---------+
          |         |         |
          v         v         v
       Worker 1  Worker 2  Worker 3
          |         |         |
          v         v         v
       results   results   results
          |         |         |
          +---------+---------+
                    |
                    v
              RESULTS CHANNEL
                    |
                    v
               MAIN GOROUTINE