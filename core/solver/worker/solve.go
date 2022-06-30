package worker

import (
	"fmt"
	"log"
	"runtime"
	"solver/core/solver"
	"solver/core/task"
	"sync"
)

// SolveWorker concrete task.Worker implementation that solves puzzles
type SolveWorker struct {
	Name          string
	TaskChannel   chan solver.Task
	ResultChannel chan task.Result
	ErrorChannel  chan error
	StopChannel   chan task.Signal
}

// NewSolveWorker return the worker that will process the solver.Tasks and put the results on the resultChannel
func NewSolveWorker(name string, taskChannel chan solver.Task, resultChannel chan task.Result, errorChannel chan error, stopChannel chan task.Signal) *SolveWorker {
	return &SolveWorker{Name: name, TaskChannel: taskChannel, ResultChannel: resultChannel, ErrorChannel: errorChannel, StopChannel: stopChannel}
}

// Solve creates SolveWorkers and initiates the solving process
func Solve[PT solver.Puzzle, ST solver.Solver](p PT, s ST) (task.Result, error) {
	// func Solve[PT puzzle.Puzzle, ST solver.Solver](p PT, s ST) (task.Result, error) {
	numCPUs := runtime.NumCPU()
	//numCPUs := 1
	// channel on which tasks can be submitted (Larger than the number of Workers)
	taskChannel := make(chan solver.Task, numCPUs*10000000)
	// channel on which the solved Board is received
	resultChannel := make(chan task.Result, 1)
	// channel on which errors can be submitted
	errorChannel := make(chan error, numCPUs+10)

	// channel to control the workers
	stopChannel := make(chan task.Signal, 1)
	defer close(resultChannel)

	// wait group which all workers will notify when done
	var wg sync.WaitGroup

	workers, err := createSolveWorkers(numCPUs, taskChannel, resultChannel, errorChannel, stopChannel)
	if err != nil {
		return nil, err
	}

	workers.Start(&wg)
	go ErrorHandler(errorChannel)

	// create the task
	tsk := solver.Task{Puzzle: p,
		Solver: s}
	taskChannel <- tsk

	// wait for a solution to arrive on the resultChannel
	solution := <-resultChannel

	close(stopChannel)  // This tells the goroutines there's nothing else to do, we have a solution
	wg.Wait()           // Wait for the goroutines to finish
	close(taskChannel)  // close the task channel
	close(errorChannel) // close the error channel

	return solution, nil
}

// createSolveWorkers create the task.Workers that will process Tasks
// provided by the tasks channel
func createSolveWorkers(numCPUs int, tasks chan solver.Task, results chan task.Result, errors chan error, stopChannel chan task.Signal) (task.Workers, error) {
	result := make(task.Workers, numCPUs)

	for i := 0; i < numCPUs; i++ {
		name := fmt.Sprintf("Worker-%02d", i)
		result[i] = NewSolveWorker(name, tasks, results, errors, stopChannel)
	}

	return result, nil
}

// Start the given task worker of tasks and when done signal completed on the given wait group
func (worker *SolveWorker) Start(wg *sync.WaitGroup) {
	for {
		// read from the tasks channel
		var t solver.Task
		select {
		case t = <-worker.TaskChannel:
			// get the concrete task
			//t = tsk.(solver.Task)
		case <-worker.StopChannel: // check if we are stopped
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
			worker.ErrorChannel <- err
		}

		// put the results returned from the solver on the results queue
		for _, r := range rs {
			worker.ResultChannel <- r
		}

		// put the puzzles returned from the solver on the task queue
		for _, p := range ps {
			tsk := solver.Task{Puzzle: p,
				Solver: t.Solver}

			worker.TaskChannel <- tsk

			// log the current size of the task queue
			noTasks := len(worker.TaskChannel)
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
