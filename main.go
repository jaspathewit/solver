package main

import "fmt"

func main() {

	dice := NewDice()

	fmt.Printf("Dice is\n%s", dice)

	clone := dice.Clone()

	fmt.Printf("Clone is \n%s", clone)

	dice = NewDice()
	dice.RollNorth()

	fmt.Printf("Dice after North is \n%s", dice)

	dice = NewDice()
	dice.RollEast()

	fmt.Printf("Dice after East is \n%s", dice)

	dice = NewDice()
	dice.RollSouth()

	fmt.Printf("Dice after South is \n%s", dice)

	dice = NewDice()
	dice.RollWest()

	fmt.Printf("Dice after West is \n%s", dice)

}
