package cmd

import (
	"github.com/spf13/cobra"
)

var inputFilename string

var rootCmd = &cobra.Command{
	Use:   "hashi",
	Short: "hashi - an application to solve hashi",
	Long:  `hashi is a general purpose hashi solver`,
	RunE:  Default,
}

func init() {
	rootCmd.Flags().StringVarP(&inputFilename, "input", "i", "", "input hashi file")
}

func Execute() error {
	return rootCmd.Execute()
}
