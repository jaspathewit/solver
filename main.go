package main

import (
	"fmt"
	"log"
	"question20/puzzle"
	"question20/solver"
	"question20/solver/worker"
	"time"
)

var _ solver.Solver = puzzle.Question20Solver{}

func main() {
	defer LogDuration(time.Now(), "question20")
	// create the starting puzzle
	p := puzzle.NewPuzzle()
	// create the solver for the puzzle
	s := puzzle.Question20Solver{}
	// start the worker.Solve with the starting puzzle and the solver
	result, err := worker.Solve(p, s)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// get the concrete type of solution
	solution := result.(puzzle.Puzzle)
	fmt.Printf("Solution is\n: %s", solution)
	fmt.Printf("Total of all cells is: %d\n", solution.Total())
}

// LogDuration logs how long a method took to execute
func LogDuration(start time.Time, name string) {
	elapsed := time.Since(start)
	if elapsed.Nanoseconds() != 0 {
		log.Printf("%s took %s", name, elapsed)
	}
}
