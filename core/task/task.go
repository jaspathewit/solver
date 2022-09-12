package task

// import (
// 	"sync"
// )

// // Workers a slice of Worker
// type Workers []Worker

// // Worker type that can be started and stopped
// type Worker interface {
// 	Start(wg *sync.WaitGroup)
// 	Stop()
// }

// // Signal used by signal channels
// type Signal struct{}

// // Start starts all the workers in a slice workers
// func (workers Workers) Start(wg *sync.WaitGroup) {
// 	for _, w := range workers {
// 		// add one to the wait group
// 		wg.Add(1)
// 		go w.Start(wg)
// 	}
// }
