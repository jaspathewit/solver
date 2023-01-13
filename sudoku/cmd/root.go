package cmd

import (
	"solver/sudoku/puzzle"

	"github.com/spf13/cobra"
)

var inputFilename string

var rootCmd = &cobra.Command{
	Use:   "sudoku",
	Short: "sudoku - an application to solve sudokus",
	Long:  `sudoku is a general purpose sudoku solver that can solve sudokus of various topologies`,
	RunE:  Default,
}

func init() {
	rootCmd.Flags().StringVarP(&inputFilename, "input", "i", "", "input sudoku file")
}

func Execute() error {
	return rootCmd.Execute()
}

// puzzelLibelle
func puzzelLibelle() (puzzle.Puzzle, error) {

	topology := puzzle.Normal{}

	g, err := puzzle.NewPuzzle(topology)
	if err != nil {
		return puzzle.Puzzle{}, err
	}

	// set up the starting values
	g.Set("1_2", "5")
	g.Set("1_3", "2")
	g.Set("1_5", "6")
	g.Set("1_6", "8")
	g.Set("1_8", "3")
	g.Set("2_2", "7")
	g.Set("2_5", "5")
	g.Set("2_7", "9")
	g.Set("2_8", "2")
	g.Set("3_2", "3")
	g.Set("3_5", "1")
	g.Set("3_9", "6")
	g.Set("5_3", "4")
	g.Set("5_6", "5")
	g.Set("5_7", "6")
	g.Set("6_3", "8")
	g.Set("6_5", "4")
	g.Set("6_7", "2")
	g.Set("7_1", "1")
	g.Set("7_2", "9")
	g.Set("7_6", "2")
	g.Set("7_8", "7")
	g.Set("8_6", "6")
	g.Set("8_9", "2")
	g.Set("9_5", "8")
	g.Set("9_7", "1")

	return g, nil
}
