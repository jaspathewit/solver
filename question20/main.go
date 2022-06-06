package main

import (
	"fmt"
	"log"
	"solver/core/solver"
	"solver/core/solver/worker"
	"solver/core/util"
	"solver/question20/puzzle"
	"time"
)

var _ solver.Solver = puzzle.Question20Solver{}

func main() {
	defer util.LogDuration(time.Now(), "question20")
	// create the starting question20
	p, err := puzzle.NewPuzzle()
	if err != nil {
		log.Fatalf("could not create puzzle")
	}
	// create the solver for the question20
	s := puzzle.Question20Solver{}
	// start the worker.Solve with the starting question20 and the solver
	result, err := worker.Solve(p, s)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// get the concrete type of solution
	solution := result.(*puzzle.Puzzle)
	fmt.Printf("Solution is\n: %s", solution)
	fmt.Printf("Total of all cells is: %d\n", solution.Total())
}
