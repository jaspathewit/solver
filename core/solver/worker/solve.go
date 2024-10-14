package worker

import (
	"fmt"
	"log"
	"runtime"
	"solver/core/solver"
	"sync"
)

// SolveWorker concrete Worker implementation that solves puzzles
type SolveWorker[PT solver.Puzzle] struct {
	Name          string
	Solver        solver.Solver[PT]
	PuzzleChannel chan PT
	ResultChannel chan PT
	ErrorChannel  chan error
	StopChannel   chan Signal
}

// NewSolveWorker return the worker that will process the Puzzles and put the results on the resultChannel
func NewSolveWorker[PT solver.Puzzle, ST solver.Solver[PT]](name string, solver ST, puzzleChannel chan PT, resultChannel chan PT, errorChannel chan error, stopChannel chan Signal) *SolveWorker[PT] {
	return &SolveWorker[PT]{Name: name,
		Solver:        solver,
		PuzzleChannel: puzzleChannel,
		ResultChannel: resultChannel,
		ErrorChannel:  errorChannel,
		StopChannel:   stopChannel}
}

// Solve creates SolveWorkers and initiates the solving process
func Solve[PT solver.Puzzle](p PT, s solver.Solver[PT]) (PT, error) {
	numCPUs := runtime.NumCPU()
	//numCPUs := 1
	// channel on which puzzles can be submitted (Larger than the number of Workers)
	puzzleChannel := make(chan PT, numCPUs*1000000)
	// channel on which the solved Puzzle is received
	resultChannel := make(chan PT, 1)
	// channel on which errors can be submitted
	errorChannel := make(chan error, numCPUs+10)
	go ErrorHandler(errorChannel)

	// channel to control the workers
	stopChannel := make(chan Signal, 1)
	defer close(resultChannel)

	// wait group which all workers will notify when done
	var wg sync.WaitGroup

	workers, err := createSolveWorkers(numCPUs, s, puzzleChannel, resultChannel, errorChannel, stopChannel)
	if err != nil {
		var result PT
		return result, err
	}

	workers.Start(&wg)

	// put the puzzle on the channel
	puzzleChannel <- p

	// wait for a result to arrive on the resultChannel
	result := <-resultChannel

	close(stopChannel)   // This tells the goroutines there's nothing else to do, we have a solution
	wg.Wait()            // Wait for the goroutines to finish
	close(puzzleChannel) // close the puzzle channel
	close(errorChannel)  // close the error channel

	return result, nil
}

// createSolveWorkers create the Workers that will process puzzles
// provided by the puzzles channel
func createSolveWorkers[PT solver.Puzzle](numCPUs int, solver solver.Solver[PT], puzzles chan PT, results chan PT, errors chan error, stopChannel chan Signal) (Workers, error) {
	result := make(Workers, numCPUs)

	for i := 0; i < numCPUs; i++ {
		name := fmt.Sprintf("Worker-%02d", i)
		result[i] = NewSolveWorker(name, solver, puzzles, results, errors, stopChannel)
	}

	return result, nil
}

// Start the given worker and when done signal completed on the given wait group
func (worker *SolveWorker[PT]) Start(wg *sync.WaitGroup) {
	// read puzzles from the puzzle channel till a limit is reached
	for {
		// read from the puzzles channel
		var p PT
		select {
		case p = <-worker.PuzzleChannel:
		case <-worker.StopChannel: // check if we are stopped
			worker.Stop()
			wg.Done()
			return
		}

		// use the Solver to solve the puzzle
		// results in a []Puzzle
		ps, rs, err := worker.Solver.Solve(p)
		if err != nil {
			// put the error why puzzle could not be solved
			// on the error channel
			err = fmt.Errorf("%s: failed solving: %s", worker.Name, err)
			worker.ErrorChannel <- err
		}

		// put the results returned from the solver on the results queue
		for _, r := range rs {
			worker.ResultChannel <- r
		}

		// put the puzzles returned from the solver on the puzzle channel
		//log.Printf("Adding %d puzzles to the channel of %d", len(ps), len(worker.PuzzleChannel))
		for _, p := range ps {
			worker.PuzzleChannel <- p

			// log the current size of the puzzle queue
			noPuzzles := len(worker.PuzzleChannel)
			if (noPuzzles % 10000) == 0 {
				log.Printf("Current number of puzzles: %d\n", noPuzzles)
				log.Printf("Last Puzzle added\n%s", p)
			}
		}
	}
}

// Stop the SolverWorker signal wait group that we are done
func (worker *SolveWorker[PT]) Stop() {
	log.Printf("Stopped worker %s\n", worker.Name)
}
