package main

import (
	"fmt"
	"log"
	"os"
	"question20/puzzle"
	"question20/solver"
	"question20/solver/worker"
)

var _ solver.Solver = puzzle.Question20Solver{}

func main() {
	// set up the log to file
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("question20.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	p := puzzle.NewPuzzle()
	s := puzzle.Question20Solver{}
	result, err := worker.Solve(p, s)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	solution := result.(puzzle.Puzzle)
	fmt.Printf("Solution is: %d", solution)

	// dice := puzzle.NewDice()

	// fmt.Printf("Dice is\n%s", dice)

	// // clone := dice.Clone()

	// // fmt.Printf("Clone is \n%s", clone)

	// dice = puzzle.NewDice()
	// dice = dice.Roll(puzzle.DirectionNorth)

	// fmt.Printf("Dice after North is \n%s", dice)

	// dice = puzzle.NewDice()
	// dice = dice.Roll(puzzle.DirectionEast)

	// fmt.Printf("Dice after East is \n%s", dice)

	// dice = puzzle.NewDice()
	// dice = dice.Roll(puzzle.DirectionSouth)

	// fmt.Printf("Dice after South is \n%s", dice)

	// dice = puzzle.NewDice()
	// dice = dice.Roll(puzzle.DirectionWest)

	// fmt.Printf("Dice after West is \n%s", dice)

	// p := puzzle.NewPuzzle()

	// fmt.Printf("Puzzle is\n%s", p)

	// solver := puzzle.Question20Solver{}

	// ps, _, _ := solver.Solve(p)

	// fmt.Printf("New Puzzles \n")

	// for _, puz := range ps {
	// 	fmt.Printf("%s\n", puz)
	// }

}
