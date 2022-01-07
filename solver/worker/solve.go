package worker

import (
	"fmt"
	"log"
	"question20/puzzle"
	"question20/solver"
	"question20/task"
	"runtime"
	"sync"
)

// SolveWorker concrete task.Worker implementation that solves puzzles
type SolveWorker struct {
	Name string
}

// Solve creates SolveWorkers and initiates the solving process
func Solve(puzzle puzzle.Puzzle, s solver.Solver) (task.Result, error) {

	numCPUs := runtime.NumCPU()
	//numCPUs := 1
	// channel on which tasks can be submitted (Larger than the number of Workers)
	taskChannel := make(chan task.Task, numCPUs*10000000)
	// channel on which errors can be submitted
	errorChannel := make(chan error, numCPUs+10)
	// channel on which the solved Board is received
	resultChannel := make(chan task.Result, 1)
	// channel to control the workers
	stopChannel := make(chan task.Signal, 1)
	defer close(resultChannel)

	// wait group which all workers will notify when done
	var wg sync.WaitGroup

	workers, err := createSolveWorkers(numCPUs)
	if err != nil {
		return nil, err
	}

	workers.Start(taskChannel, resultChannel, errorChannel, stopChannel, &wg)
	go ErrorHandler(errorChannel)

	// create the first and only puzzle
	tsk := solver.Task{Puzzle: puzzle,
		Solver: s}
	taskChannel <- tsk

	// wait for a solution to arrive on the resultChannel
	solution := <-resultChannel

	close(stopChannel)  // This tells the goroutines there's nothing else to do, we have a solution
	wg.Wait()           // Wait for the goroutines to finish
	close(taskChannel) // close the task channel
	close(errorChannel) // close the error channel

	return solution, nil
}

// createSolveWorkers create the task.Workers that will process Tasks
// by solving the puzzle
func createSolveWorkers(numCPUs int) (task.Workers, error) {
	result := make(task.Workers, numCPUs)

	for i := 0; i < numCPUs; i++ {
		name := fmt.Sprintf("Worker-%02d", i)
		result[i] = &SolveWorker{Name: name}
	}

	return result, nil
}

// Start the given task worker of tasks and when done signal completed on the given wait group
func (worker *SolveWorker) Start(tasks chan task.Task, results chan task.Result, errors chan error, stopChannel chan task.Signal, wg *sync.WaitGroup) {
	for {
		// read from the tasks channel
		var t solver.Task
		select {
		case tsk := <-tasks:
			// get the concrete task
			t = tsk.(solver.Task)
		case <-stopChannel: // check if we are stopped
			worker.Stop()
			wg.Done()
			return
		}

		// use the Solver to solve the puzzle in the task
		// results in a []Puzzle
		ps, rs, err := t.Solver.Solve(t.Puzzle)
		if err != nil {
			// log the reason that the board could not be solved
			err = fmt.Errorf("%s: failed solving: %s", worker.Name, err)
			errors <- err
		}

		// put the results returned from the solver on the results queue
		for _, r := range rs {
			results <- r
		}

		// put the puzzles returned from the solver on the task queue
		for _, p := range ps {
			tsk := solver.Task{Puzzle: p,
				Solver: t.Solver}

			tasks <- tsk

			// log the current size of the task queue
			noTasks := len(tasks)
			if (noTasks % 10000) == 0 {
				log.Printf("Current number of tasks: %d\n", noTasks)
				log.Printf("Last Puzzle added\n%s", p)
			}
		}
	}
}

// Stop the SolverWorker signal wait group that we are done
func (worker *SolveWorker) Stop() {
	log.Printf("Stopped worker %s\n", worker.Name)
}
