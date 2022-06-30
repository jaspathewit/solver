package worker

import (
	"fmt"
	"log"
	"solver/core/solver"
	"solver/core/task"
	"sync"
)

// TaskWorker concrete worker.Worker implementation
// receives "tasks" from a task Channel and processe them
type TaskWorker[T Task] struct {
	Name          string
	TaskChannel   chan T
	ResultChannel chan Result
	ErrorChannel  chan error
	StopChannel   chan Signal
}

// NewTaskWorker return the worker that will process tasks and put the results on the resultChannel
func NewTaskWorker[T Task](name string, taskChannel chan T, resultChannel chan Result, errorChannel chan error, stopChannel chan Signal) *TaskWorker[T] {
	return &TaskWorker[T]{Name: name, TaskChannel: taskChannel, ResultChannel: resultChannel, ErrorChannel: errorChannel, StopChannel: stopChannel}
}

// CreateWorkers create the TaskWorkers that will process Tasks
// provided by the tasks channel
func CreateWorkers[T Task](numCPUs int, tasks chan T, results chan Result, errors chan error, stopChannel chan Signal) (task.Workers, error) {
	result := make(task.Workers, numCPUs)

	for i := 0; i < numCPUs; i++ {
		name := fmt.Sprintf("Worker-%02d", i)
		result[i] = NewTaskWorker[T](name, tasks, results, errors, stopChannel)
	}

	return result, nil
}

// Start the given TaskWorker of tasks and when done signal completed on the given wait group
func (worker *TaskWorker[T]) Start(wg *sync.WaitGroup) {
	for {
		// read from the tasks channel
		var t T
		select {
		case t = <-worker.TaskChannel:
			// get the concrete task
			//t = tsk.(solver.Task)
		case <-worker.StopChannel: // check if we are stopped
			worker.Stop()
			wg.Done()
			return
		}

		// process the task
		ps, rs, err := t.Process() // Solver.Solve(t.Puzzle)
		if err != nil {
			// log the error
			err = fmt.Errorf("%s: failed processing: %s", worker.Name, err)
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
func (worker *TaskWorker[T]) Stop() {
	log.Printf("Stopped worker %s\n", worker.Name)
}
