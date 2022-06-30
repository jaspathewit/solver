package worker

import (
    "sync"
)

// Workers a slice of Worker
type Workers []Worker

// Worker type that can be started and stopped
type Worker interface {
    Start(wg *sync.WaitGroup)
    Stop()
}

// Task interface implemented by tasks that can be submitted to a task channel to
type Task interface {
    Process() (
}

// Result interface implemented by results that can be submitted to a result channel
type Result interface {
}

// Signal used by signal channels
type Signal struct{}

//Start starts a number of workers processing a channel of tasks submitted on the given task channel
func (workers Workers) Start(wg *sync.WaitGroup) {
    for _, w := range workers {
        // add one to the wait group
        wg.Add(1)
        go w.Start(wg)
    }
}
