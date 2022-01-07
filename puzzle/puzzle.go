package puzzle

import (
	"bytes"
)

type Puzzles []Puzzle

type Puzzle struct {
	Board  []Cells
	Labels Labels

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
	for _, row := range p.Board[2:10] {
		buffer.WriteString(row[2:10].String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// Total the total of all the Cells in the Puzzle
func (p Puzzle) Total() int64 {
	var result int64
	for _, row := range p.Board[2:10] {
		for _, cell := range row[2:10] {
			result += int64(cell.Value)
		}
	}
	return result
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

	// roll the dice in the that direction
	dice := c.Dice.Roll(direction)

	// check if the roll was valid direction (
	if c.Board[dice.Row][dice.Col].Value != 0 {
		return Puzzle{}, false
	}

	c.Dice = dice
	c.Board[c.Dice.Row][c.Dice.Col].Value = c.Dice.Top

	// check if the after the roll of the dice the puzzle is still valid
	if !c.Valid() {
		return Puzzle{}, false
	}

	// check if after the roll of the dice the puzzle is still solvable
	if c.Partitioned() {
		return Puzzle{}, false
	}

	return c, true
}

// Solved checks if the given puzzle is solved
func (p Puzzle) Solved() bool {
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

// Valid checks if the given puzzle is valid
// the current value of the cell containing the dice
// must not be the same as any cell reachable from current
// dice location via the knights move
// .......
// ..0.0..
// .0...0.
// ...X...
// .0...0.
// ..0.0..
// .......
func (p Puzzle) Valid() bool {
	value := p.Board[p.Dice.Row][p.Dice.Col].Value

	if value == p.Board[p.Dice.Row-2][p.Dice.Col-1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row-2][p.Dice.Col+1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row-1][p.Dice.Col+2].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+1][p.Dice.Col+2].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+2][p.Dice.Col+1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+2][p.Dice.Col-1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+1][p.Dice.Col-2].Value {
		return false
	}

	if value == p.Board[p.Dice.Row-1][p.Dice.Col-2].Value {
		return false
	}

	return true
}

// Partitioned determins if all the unassigned cells are still all four way connected
func (p Puzzle) Partitioned() bool {
	// create the labels for the cells and add them to the puzzle
	p.Labels = NewLabels(len(p.Board))
	defer func() { p.Labels = nil }()

	currentLabel := int8(0)
	// apply the ccl algorithum
	for row := range p.Board {
		for col := range p.Board[row] {
			// check if this cell is not labeled and does not have a value
			if p.Labels[row][col] == 0 && p.Board[row][col].Value == 0 {
				// add one to the current label
				currentLabel += 1

				// check if we have found a new partition
				if currentLabel >= 2 {
					return true
				}

				p.dfs(row, col, currentLabel)
			}
		}
	}

	//if currentLabel >= 2 {
	//	log.Printf("Puzzle is partitioned\n%s", p.Labels.String())
	//	log.Printf("Puzzle\n%s", p.String())
	//	return true
	//}

	return false
}

// dfs perform depth first search on the puzzle
func (p Puzzle) dfs(row int, col int, currentLabel int8) {
	size := len(p.Board)
	// check if we are move outside the bounds of the board
	if row < 0 || row == size {
		return // out of bounds
	}

	if col < 0 || col == size {
		return // out of bounds
	}

	// Check that the cell under consideration is labeled or not visited yet
	if p.Labels[row][col] != 0 || p.Board[row][col].Value != 0 {
		return // already labeled or not marked with 0 in m
	}

	// mark the current cell
	p.Labels[row][col] = currentLabel

	// recursively mark the 4 neighbors
	p.dfs(row-1, col, currentLabel)
	p.dfs(row, col+1, currentLabel)
	p.dfs(row+1, col, currentLabel)
	p.dfs(row, col-1, currentLabel)
}
