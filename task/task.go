package task

import (
	"sync"
)

// Workers a slice of Worker
type Workers []Worker

// Worker type that will process a channel of tasks
type Worker interface {
	Start(tasks chan Task, results chan Result, errors chan error, wg *sync.WaitGroup)
	Stop()
}

// Task interface implemented by tasks that can be submitted to a task channel to
// be executed by a number of worker routiens.
type Task interface {
}

// Result interface implemented by results that can be submitted to a result channel
type Result interface {
}

//Start starts a number of workers processing a channel of tasks submitted on the given task channel
func (workers Workers) Start(tasks chan Task, results chan Result, errors chan error, wg *sync.WaitGroup) {
	for _, w := range workers {
		// add one to the wait group
		wg.Add(1)
		go w.Start(tasks, results, errors, wg)
	}
}
