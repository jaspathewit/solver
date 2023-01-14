package cmd

import (
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
