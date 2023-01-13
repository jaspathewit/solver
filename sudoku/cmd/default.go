package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"solver/core/solver/worker"
	"solver/sudoku/puzzle"

	"github.com/spf13/cobra"
)

var mapTopologyTypeToTopology = map[TopologyType]puzzle.Topology{TopologyTypeNormal: puzzle.Normal{}, TopologyTypeSamuri: puzzle.Samuri{}}

// Default command loads the sudoku from a file and solves it
func Default(cmd *cobra.Command, args []string) error {
	p, err := loadSudokuFromFile(inputFilename)
	if err != nil {
		return fmt.Errorf("failed to load puzzel: %s", err)
	}

	// create the solver for the suduku
	s := puzzle.SudokuSolver[puzzle.Puzzle]{}
	// start the worker.Solve with the starting sudoku and the solver
	result, err := worker.Solve[puzzle.Puzzle](p, s)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	fmt.Printf("Solution is:\n%s", result)
	return nil
}

func loadSudokuFromFile(filename string) (puzzle.Puzzle, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return puzzle.Puzzle{}, fmt.Errorf("could not load file [%s]: %s", filename, err)
	}

	s := Sudoku{}

	err = xml.Unmarshal([]byte(data), &s)
	if err != nil {
		return puzzle.Puzzle{}, fmt.Errorf("could not load file [%s]: %s", filename, err)
	}

	topology, err := getTopology(s.Topology)
	if err != nil {
		return puzzle.Puzzle{}, fmt.Errorf("could not load file [%s]: %w", err)
	}

	result, err := puzzle.NewPuzzle(topology)
	if err != nil {
		return puzzle.Puzzle{}, err
	}

	// fill the cells
	for i, row := range s.Rows {
		for j, r := range row {
			// convert the rune to a string
			c := string(r)

			// if there is a value in the cell
			if c != "." {
				// create the cell reference
				cellRef := fmt.Sprintf("%d_%d", i+1, j+1)
				result.Set(cellRef, c)
			}
		}
	}

	return result, err

}

func getTopology(topology TopologyType) (puzzle.Topology, error) {
	// get the topology
	result, ok := mapTopologyTypeToTopology[topology]
	if !ok {
		return result, fmt.Errorf("topology %s not supported", topology)
	}
	return result, nil
}
