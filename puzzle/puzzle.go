package puzzle

import "bytes"

type Puzzles []Puzzle

type Puzzle struct {
	Board []Cells

	Dice Dice
}

// NewPuzzel constructs a new puzzle.Puzzle
func NewPuzzle() Puzzle {
	result := Puzzle{}
	result.Dice = NewDice()

	// put the dice in the starting position
	result.Dice.Row = 3
	result.Dice.Col = 5

	result.Board = make([]Cells, 12)
	for i := range result.Board {
		result.Board[i] = make(Cells, 12)
		result.Board[i][0].Value = -1
		result.Board[i][1].Value = -1
		result.Board[i][10].Value = -1
		result.Board[i][11].Value = -1

		if i == 0 || i == 1 || i == 10 || i == 11 {
			for j := range result.Board[i] {
				result.Board[i][j].Value = -1
			}
		}
	}

	result.Board[result.Dice.Row][result.Dice.Col].Value = result.Dice.Top

	return result
}

// Clone constructs a clone of the puzzle.Puzzle
func (p Puzzle) Clone() Puzzle {

	result := NewPuzzle()
	// clone the dice
	result.Dice = p.Dice

	for row := range result.Board {
		copy(result.Board[row], p.Board[row])
	}

	return result
}

// String representation of the Puzzle
func (p Puzzle) String() string {
	var buffer bytes.Buffer
	for _, cells := range p.Board[2:10] {
		buffer.WriteString(cells[2:10].String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// Move "moves" the puzzle in the given direction
// returns the puzzle representation after rolling the dice
// in the direction given, and ok if the Move in the the
// given direction is possible.
func (p Puzzle) Move(direction Direction) (Puzzle, bool) {
	// take a clone of the given Puzzle
	c := p.Clone()

	// record that the dice is moving in the given direction from the current cell
	c.Board[c.Dice.Row][c.Dice.Col].Direction = direction

	// roll the dice in the tha direction
	dice := c.Dice.Roll(direction)

	// check if the roll was valid direction
	if c.Board[dice.Row][dice.Col].Value != 0 {
		return Puzzle{}, false
	}

	c.Dice = dice
	c.Board[c.Dice.Row][c.Dice.Col].Value = c.Dice.Top

	return c, true
}

// IsSolved checks if the given puzzle is solved
func (p Puzzle) IsSolved() bool {
	result := p.isDiceLocationCorrect() && p.noZeroCells()

	return result

}

// isDiceLocation checks if the dice is in one of the correct locations
func (p Puzzle) isDiceLocationCorrect() bool {
	if p.Dice.Row == 2 && p.Dice.Col == 3 {
		return true
	}

	if p.Dice.Row == 2 && p.Dice.Col == 7 {
		return true
	}

	if p.Dice.Row == 4 && p.Dice.Col == 3 {
		return true
	}

	if p.Dice.Row == 4 && p.Dice.Col == 7 {
		return true
	}

	if p.Dice.Row == 5 && p.Dice.Col == 4 {
		return true
	}

	if p.Dice.Row == 5 && p.Dice.Col == 6 {
		return true
	}

	return false
}

// noZeroCells checks if there are any zero valued cells
func (p Puzzle) noZeroCells() bool {
	size := len(p.Board)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if p.Board[row][col].Value == 0 {
				return false
			}
		}
	}

	return true
}
