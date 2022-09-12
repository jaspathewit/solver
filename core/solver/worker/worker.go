package worker

import (
	"log"
	"sync"
)

// Workers a slice of Worker
type Workers []Worker

// Worker type that can be started and stopped
type Worker interface {
	Start(wg *sync.WaitGroup)
	Stop()
}

// Signal used by signal channels
type Signal struct{}

// Start starts all the workers in a slice workers
func (workers Workers) Start(wg *sync.WaitGroup) {
	for _, w := range workers {
		// add one to the wait group
		wg.Add(1)
		go w.Start(wg)
	}
}

// ErrorHandler Handles the errors sent on an error channel
// when the error channel is closed it sends the total number
// of errors handled on the given totalErrors channel
func ErrorHandler(errors chan error) {
	// process the errors on the channel
	for {
		// read from the errors channel
		err, ok := <-errors

		// test if the channel has been closed
		if !ok {
			return
		}

		log.Print(err)
	}
}
