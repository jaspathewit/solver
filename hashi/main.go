package main

import (
	"log"
	"solver/core/solver"
	"solver/core/util"
	"solver/hashi/cmd"
	"solver/hashi/puzzle"
	"time"
)

var _ solver.Solver[puzzle.Puzzle] = puzzle.HashiSolver[puzzle.Puzzle]{}

func main() {
	defer util.LogDuration(time.Now(), "hashi")

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
