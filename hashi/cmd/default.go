package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"solver/core/solver/worker"
	"solver/hashi/puzzle"

	"github.com/spf13/cobra"
)

var mapTopologyTypeToTopology = map[TopologyType]puzzle.Topology{TopologyTypeNormal: puzzle.Normal{}}

// Default command loads the sudoku from a file and solves it
func Default(cmd *cobra.Command, args []string) error {
	p, err := loadHashiFromFile(inputFilename)
	if err != nil {
		return fmt.Errorf("failed to load puzzel: %s", err)
	}

	// create the solver for the suduku
	s := puzzle.HashiSolver[puzzle.Puzzle]{}
	// start the worker.Solve with the starting sudoku and the solver
	result, err := worker.Solve[puzzle.Puzzle](p, s)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	fmt.Printf("Solution is:\n%s", result)
	return nil
}

func loadHashiFromFile(filename string) (puzzle.Puzzle, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return puzzle.Puzzle{}, fmt.Errorf("could not load file [%s]: %s", filename, err)
	}

	s := Hashi{}

	err = xml.Unmarshal([]byte(data), &s)
	if err != nil {
		return puzzle.Puzzle{}, fmt.Errorf("could not load file [%s]: %s", filename, err)
	}

	topology, err := getTopology(s.Topology)
	if err != nil {
		return puzzle.Puzzle{}, fmt.Errorf("could not load file [%s]: %s", filename, err)
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
