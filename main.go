package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"question20/puzzle"
	"question20/solver"
	"question20/solver/worker"
)

var _ solver.Solver = puzzle.Question20Solver{}

func main() {
	board := puzzle.NewBoard()
	s := puzzle.Question20Solver{}
	result, err := worker.Solve(board, s)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	solution := result.(puzzle.Board)
	fmt.Printf("Solution is: %d", solution)

	//dice := puzzle.NewDice()
	//
	//fmt.Printf("Dice is\n%s", dice)
	//
	//clone := dice.Clone()
	//
	//fmt.Printf("Clone is \n%s", clone)
	//
	//dice = puzzle.NewDice()
	//dice.RollNorth()
	//
	//fmt.Printf("Dice after North is \n%s", dice)
	//
	//dice = puzzle.NewDice()
	//dice.RollEast()
	//
	//fmt.Printf("Dice after East is \n%s", dice)
	//
	//dice = puzzle.NewDice()
	//dice.RollSouth()
	//
	//fmt.Printf("Dice after South is \n%s", dice)
	//
	//dice = puzzle.NewDice()
	//dice.RollWest()
	//
	//fmt.Printf("Dice after West is \n%s", dice)

}
