package main

import "fmt"

func main() {

	dice := NewDice()

	fmt.Printf("Dice is\n%s", dice)

	clone := dice.Clone()

	fmt.Printf("Clone is \n%s", clone)

	clone.RollNorth()

	fmt.Printf("Clone after North is \n%s", clone)

}
