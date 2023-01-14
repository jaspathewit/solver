package main

import (
	"log"
	"solver/core/solver"
	"solver/core/util"
	"solver/sudoku/cmd"
	"solver/sudoku/puzzle"
	"time"
)

var _ solver.Solver[puzzle.Puzzle] = puzzle.SudokuSolver[puzzle.Puzzle]{}

func main() {
	defer util.LogDuration(time.Now(), "sudoku")

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
